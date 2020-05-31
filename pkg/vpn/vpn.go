package vpn

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"path"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/event"
	"github.com/MemeLabs/go-ppspp/pkg/kademlia"
	"github.com/MemeLabs/go-ppspp/pkg/logutil"
	"github.com/MemeLabs/go-ppspp/pkg/pb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

const reservedPortCount uint16 = 1000

// default service ports
const (
	HashTablePort uint16 = iota + 10
	PeerIndexPort
	PeerExchangePort
	DirectoryPort
	SwarmServicePort
)

var (
	linksActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "strims_vpn_links_active",
		Help: "The number of active network links",
	})
	linkReadBytes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "strims_vpn_link_read_bytes",
		Help: "The total number of bytes read from network links",
	})
	linkWriteBytes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "strims_vpn_link_write_bytes",
		Help: "The total number of bytes written to network links",
	})
	dialCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "strims_vpn_dial_count",
		Help: "The total number of dialed network connections",
	}, []string{"scheme"})
	dialErrorCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "strims_vpn_dial_error_count",
		Help: "The total number of network connection dial errors",
	}, []string{"scheme"})
)

// HostOption ...
type HostOption func(h *Host) error

// NewHost ...
func NewHost(logger *zap.Logger, key *pb.Key, options ...HostOption) (*Host, error) {
	discriminator, err := randUint16(math.MaxUint16)
	if err != nil {
		return nil, err
	}

	id, err := NewHostID(key.Public, discriminator)
	if err != nil {
		return nil, err
	}

	h := &Host{
		logger:        logger,
		discriminator: discriminator,
		id:            id,
		key:           key,
	}

	for _, o := range options {
		if err := o(h); err != nil {
			return nil, err
		}
	}

	for _, iface := range h.interfaces {
		if listener, ok := iface.(Listener); ok {
			go listener.Listen(h)
		}
	}

	return h, nil
}

// Host ...
type Host struct {
	logger        *zap.Logger
	discriminator uint16
	id            kademlia.ID
	key           *pb.Key
	interfaces    []Interface
	peerObservers event.Observable
	peers         sync.Map

	// TODO: find a better place for this...
	networkBroker NetworkBroker
}

// Close ...
func (h *Host) Close() {
	for _, iface := range h.interfaces {
		if listener, ok := iface.(Listener); ok {
			listener.Close()
		}
	}

	h.peers.Range(func(pi, _ interface{}) bool {
		pi.(*Peer).Close()
		return true
	})
}

// Logger ...
func (h *Host) Logger() *zap.Logger {
	return h.logger
}

// Discriminator ...
func (h *Host) Discriminator() uint16 {
	return h.discriminator
}

// ID ...
func (h *Host) ID() kademlia.ID {
	return h.id
}

// Key ...
func (h *Host) Key() *pb.Key {
	return h.key
}

// AddLink ...
func (h *Host) AddLink(c Link) {
	go func() {
		linksActive.Inc()
		defer linksActive.Dec()

		p, err := newPeer(h.logger, instrumentLink(c), h.key)
		if err != nil {
			h.logger.Error("peer init error", zap.Error(err))
			return
		}

		h.logger.Debug(
			"created peer",
			logutil.ByteHex("peer", p.Certificate.Key),
			zap.String("type", reflect.TypeOf(c).String()),
			zap.Int("mtu", c.MTU()),
		)

		h.peerObservers.Emit(p)
		h.peers.Store(p, struct{}{})

		p.run()

		h.peers.Delete(p)
	}()
}

// NotifyPeer ...
func (h *Host) NotifyPeer(ch chan *Peer) {
	h.peerObservers.Notify(ch)
}

func (h *Host) dialer(scheme string) Interface {
	for _, i := range h.interfaces {
		if i.ValidScheme(scheme) {
			return i
		}
	}
	return nil
}

// Dial ...
func (h *Host) Dial(addr InterfaceAddr) error {
	scheme := addr.Scheme()

	dialCount.WithLabelValues(scheme).Inc()
	h.logger.Debug(
		"dialing",
		zap.String("scheme", scheme),
		zap.Stringer("addr", addr.(fmt.Stringer)),
	)

	d := h.dialer(scheme)
	if d == nil {
		dialErrorCount.WithLabelValues(scheme).Inc()
		return errors.New("unsupported scheme")
	}

	if err := d.Dial(h, addr); err != nil {
		dialErrorCount.WithLabelValues(scheme).Inc()
		h.logger.Error("dial error", zap.Error(err))
		return err
	}
	return nil
}

// NewHostID ...
func NewHostID(key []byte, hid uint16) (kademlia.ID, error) {
	var t [20]byte
	copy(t[:18], key)
	binary.BigEndian.PutUint16(t[18:], hid)
	return kademlia.UnmarshalID(t[:])
}

// Listener ...
type Listener interface {
	Listen(h *Host) error
	Close() error
}

// WithInterface ...
func WithInterface(i Interface) HostOption {
	return func(host *Host) error {
		host.interfaces = append(host.interfaces, i)
		return nil
	}
}

// Interface ...
type Interface interface {
	ValidScheme(string) bool
	Dial(h *Host, addr InterfaceAddr) error
}

// InterfaceAddr ...
type InterfaceAddr interface {
	Scheme() string
}

// Link ...
type Link interface {
	io.ReadWriteCloser
	MTU() int
}

func instrumentLink(l Link) *instrumentedLink {
	return &instrumentedLink{l}
}

type instrumentedLink struct {
	Link
}

func (l *instrumentedLink) Read(p []byte) (int, error) {
	n, err := l.Link.Read(p)
	linkReadBytes.Add(float64(n))
	return n, err
}

func (l *instrumentedLink) Write(p []byte) (int, error) {
	n, err := l.Link.Write(p)
	linkWriteBytes.Add(float64(n))
	return n, err
}

// WithBootstrapClients ...
func WithBootstrapClients(clients []*pb.BootstrapClient) HostOption {
	return func(host *Host) error {
		for _, o := range clients {
			switch o := o.ClientOptions.(type) {
			case *pb.BootstrapClient_WebsocketOptions:
				go host.Dial(WebSocketAddr(o.WebsocketOptions.Url))
			}
		}
		return nil
	}
}

func jsonDump(i interface{}) {
	_, file, line, _ := runtime.Caller(1)
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"%s %s:%d: %s\n",
		time.Now().Format("2006/01/02 15:04:05.000000"),
		path.Base(file),
		line, string(b),
	)
}