package service

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/MemeLabs/go-ppspp/pkg/dao"
	"github.com/MemeLabs/go-ppspp/pkg/kademlia"
	"github.com/MemeLabs/go-ppspp/pkg/pb"
	"github.com/MemeLabs/go-ppspp/pkg/ppspp"
	"github.com/MemeLabs/go-ppspp/pkg/vpn"
	"github.com/petar/GoLLRB/llrb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var (
	defaultNameChangeQuota uint32 = 5

	ErrNotFound             = errors.New("Could not find item matching criteria")
	ErrNameChangesExhausted = errors.New("no name changes remaining")
	ErrNameAlreadyTaken     = errors.New("the requested name is already in use")
	ErrUnimplemented        = errors.New("not implemented")
	ErrAlreadyExists        = errors.New("a record for this key already exists")
)

const clientTimeout = 10 * time.Second

// TODO: store in kv store
// TODO: logging
// TODO: tests
// TODO: probably more

func NewNickServ(svc *NetworkServices, cfg *pb.ServerConfig, salt []byte) (*NickServ, error) {
	key := &pb.Key{}
	if err := json.Unmarshal(cfg.Key, &key); err != nil {
		return nil, err
	}

	w, err := ppspp.NewWriter(ppspp.WriterOptions{
		// SwarmOptions: ppspp.NewDefaultSwarmOptions(),
		SwarmOptions: ppspp.SwarmOptions{
			LiveWindow: 1 << 10, // 1MB
			ChunkSize:  128,
		},
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	svc.Swarms.OpenSwarm(w.Swarm())

	port, err := svc.Network.ReservePort()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	newSwarmPeerManager(ctx, svc, getPeersGetter(ctx, svc, key.Public, salt))

	b, err := proto.Marshal(&pb.NetworkAddress{
		HostId: svc.Host.ID().Bytes(nil),
		Port:   uint32(port),
	})
	if err != nil {
		cancel()
		return nil, err
	}
	_, err = svc.HashTable.Set(ctx, key, salt, b)
	if err != nil {
		cancel()
		return nil, err
	}

	s := &NickServ{
		logger:          svc.Host.Logger(),
		close:           cancel,
		swarm:           w.Swarm(),
		svc:             svc,
		roles:           cfg.GetRoles(),
		tokenttl:        cfg.GetTokenTtl(),
		nameChangeQuota: cfg.GetNameChangeQuota(),
		store: &NickServStore{
			nicks:   make(map[string]*nickServItem),
			records: llrb.New(),
		},
	}

	// TODO: populate store from kv

	err = svc.Network.SetHandler(port, s)
	if err != nil {
		cancel()
		return nil, err
	}

	return s, nil
}

type NickServ struct {
	//	cfg             *pb.ServConfig
	logger          *zap.Logger
	close           context.CancelFunc
	swarm           *ppspp.Swarm
	closeOnce       sync.Once
	svc             *NetworkServices
	store           *NickServStore
	roles           []string
	tokenttl        uint64
	nameChangeQuota uint32 // smaller?
}

type NickServStore struct {
	// public key -> record
	records *llrb.LLRB
	// nick -> record
	nicks map[string]*nickServItem
	sync.Mutex
}

func (s *NickServ) Close() {
	s.closeOnce.Do(func() {
		s.close()
		s.svc.Swarms.CloseSwarm(s.swarm.ID())
	})
}

func (s *NickServ) HandleMessage(msg *vpn.Message) (forward bool, err error) {
	var m pb.NickServRPCCommand
	if err := proto.Unmarshal(msg.Body, &m); err != nil {
		return true, err
	}

	valid := ed25519.Verify(m.SourcePublicKey, msg.Body, msg.Trailers[0].Signature)
	if !valid {
		return false, dao.ErrInvalidSignature
	}

	var resp proto.Message

	switch b := m.Body.(type) {
	case *pb.NickServRPCCommand_Create_:
		resp, err = s.handleCreate(m.SourcePublicKey, b.Create)
	case *pb.NickServRPCCommand_Retrieve_:
		resp, err = s.handleRetrieve(m.SourcePublicKey)
	case *pb.NickServRPCCommand_Update_:
		resp, err = s.handleUpdate(m.SourcePublicKey, b.Update)
	case *pb.NickServRPCCommand_Delete_:
		resp, err = s.handleDelete(m.SourcePublicKey)
	default:
		err = errors.New("unexpected message type")
	}

	if err != nil {
		resp = &pb.NickServRPCResponse{
			Body: &pb.NickServRPCResponse_Error{
				Error: err.Error(),
			},
		}
	}

	// TODO: return some errors that can occur during handling

	return false, s.Send(resp, msg.Trailers[0].HostID, msg.Header.SrcPort, msg.Header.DstPort)
}

func (s *NickServ) Send(msg proto.Message, dstID kademlia.ID, dstPort, srcPort uint16) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return s.svc.Network.Send(dstID, dstPort, srcPort, msgBytes)
}

// serializeNickToken returns a stable byte representation of a NickServToken
func serializeNickToken(token *pb.NickServToken) ([]byte, int) {
	var rolesLength int
	for _, role := range token.Roles {
		rolesLength += len(role)
	}

	b := make([]byte, len(token.Key)+len(token.Nick)+rolesLength+8)

	n := copy(b, token.Key)
	n += copy(b[n:], []byte(token.Nick))
	for _, role := range token.Roles {
		n += copy(b[n:], []byte(role))
	}
	binary.BigEndian.PutUint64(b[n:], token.ValidUntil)
	n += 8

	return b, n
}

func (s *NickServ) handleCreate(key []byte, msg *pb.NickServRPCCommand_Create) (*pb.NickServRPCResponse, error) {
	_, ok := s.store.nicks[strings.ToLower(msg.Nick)]
	if ok {
		return nil, ErrNameAlreadyTaken
	}

	now := time.Now().UTC()
	rid, err := dao.GenerateSnowflake()
	if err != nil {
		return nil, err
	}
	r := &pb.NickServRecord{
		Id:                       rid,
		Key:                      key,
		Nick:                     msg.Nick,
		CreatedTimestamp:         uint64(now.Unix()),
		UpdatedTimestamp:         uint64(now.Unix()),
		RemainingNameChangeQuota: defaultNameChangeQuota,
		Roles:                    []string{},
	}

	err = s.store.Insert(r)
	if err != nil {
		return nil, err
	}

	token := &pb.NickServToken{
		Key:  r.Key,
		Nick: r.Nick,
		//Roles:      r.Roles, TODO: handle roles
		ValidUntil: s.validUntil(),
	}

	err = signNickToken(token, s.svc.Host.Key())
	if err != nil {
		return nil, err
	}

	resp := &pb.NickServRPCResponse{
		Body: &pb.NickServRPCResponse_Create{
			Create: token,
		},
	}

	return resp, nil
}

func (s *NickServ) handleRetrieve(key []byte) (*pb.NickServRPCResponse, error) {
	item, err := s.store.Retrieve(key)
	if err != nil {
		return nil, err
	}

	token := &pb.NickServToken{
		Key:  item.Key,
		Nick: item.Nick,
		//Roles:      item.Roles, TODO: handle roles
		ValidUntil: s.validUntil(),
	}

	err = signNickToken(token, s.svc.Host.Key())
	if err != nil {
		return nil, err
	}

	body := &pb.NickServRPCResponse_Retrieve{
		Retrieve: token,
	}

	return &pb.NickServRPCResponse{
		Body: body,
	}, nil
}

func (s *NickServ) handleUpdate(key []byte, msg *pb.NickServRPCCommand_Update) (*pb.NickServRPCResponse, error) {
	record, err := s.store.Retrieve(key)
	if err != nil {
		return nil, err
	}

	switch t := msg.Param.(type) {
	case *pb.NickServRPCCommand_Update_NameChangeQuota:
		// TODO handle permissions for quota changes
		return nil, ErrUnimplemented
	case *pb.NickServRPCCommand_Update_Nick:
		if record.RemainingNameChangeQuota <= 0 {
			return nil, ErrNameChangesExhausted
		}

		if record.Nick != t.Nick.OldNick {
			return nil, errors.New("nick doesn't match record")
		}

		_, ok := s.store.nicks[strings.ToLower(t.Nick.NewNick)]
		if ok {
			return nil, ErrNameAlreadyTaken
		}

		// TODO: should we create a new record instead of modifying the old one?
		record.Nick = t.Nick.NewNick
		record.RemainingNameChangeQuota--
		err = s.store.Update(record, t.Nick.OldNick)
		if err != nil {
			return nil, err
		}
	case *pb.NickServRPCCommand_Update_Roles_:
		// TODO handle permissions for role changes
		return nil, ErrUnimplemented
	}

	record.UpdatedTimestamp = uint64(time.Now().UTC().Unix())

	return &pb.NickServRPCResponse{
		Body: &pb.NickServRPCResponse_Update_{
			Update: &pb.NickServRPCResponse_Update{},
		},
	}, nil
}

func (s *NickServ) handleDelete(key []byte) (*pb.NickServRPCResponse, error) {
	return &pb.NickServRPCResponse{
		Body: &pb.NickServRPCResponse_Delete_{},
	}, s.store.Delete(key)
}

func NewNickServClient(svc *NetworkServices, key, salt []byte) (*NickServClient, error) {
	port, err := svc.Network.ReservePort()
	if err != nil {
		return nil, err
	}

	swarm, err := ppspp.NewSwarm(
		ppspp.NewSwarmID(key),
		// ppspp.NewDefaultSwarmOptions(),
		ppspp.SwarmOptions{
			LiveWindow: 1 << 10, // 1MB
			ChunkSize:  128,
		},
	)
	if err != nil {
		return nil, err
	}
	svc.Swarms.OpenSwarm(swarm)

	ctx, cancel := context.WithCancel(context.Background())

	err = svc.PeerIndex.Publish(ctx, key, salt, 0)
	if err != nil {
		svc.Swarms.CloseSwarm(swarm.ID())
		cancel()
		return nil, err
	}

	newSwarmPeerManager(ctx, svc, getPeersGetter(ctx, svc, key, salt))

	c := &NickServClient{
		close:     cancel,
		svc:       svc,
		port:      port,
		responses: make(map[uint64]chan *pb.NickServRPCResponse),
	}

	go c.syncAddr(ctx, svc, key, salt)

	return nil, nil
}

type NickServClient struct {
	close     context.CancelFunc
	closeOnce sync.Once
	svc       *NetworkServices
	responses map[uint64]chan *pb.NickServRPCResponse
	hostAddr  atomic.Value
	port      uint16
}

func (c *NickServClient) syncAddr(ctx context.Context, svc *NetworkServices, key, salt []byte) {
	var nextTick time.Duration
	for {
		select {
		case <-time.After(nextTick):
			addr, err := getHostAddr(ctx, svc, key, salt)
			if err != nil {
				nextTick = syncAddrRetryIvl
				continue
			}

			c.hostAddr.Store(addr)

			nextTick = syncAddrRefreshIvl
		case <-ctx.Done():
			return
		}
	}
}

// Close ...
func (c *NickServClient) Close() {
	c.closeOnce.Do(func() {
		c.close()
	})
}

func (c *NickServClient) MessageHandler(msg *pb.NickServRPCResponse) {
	ch, ok := c.responses[msg.RequestId]
	if !ok {
		return
	}
	select {
	case ch <- msg:
	default:
	}
}

func (c *NickServClient) registerResponseChan() (uint64, error) {
	rid, err := dao.GenerateSnowflake()
	if err != nil {
		return 0, err
	}

	c.responses[rid] = make(chan *pb.NickServRPCResponse)

	return rid, nil
}

func (c *NickServClient) awaitResponse(ctx context.Context, rid uint64) (*pb.NickServRPCResponse, error) {
	defer delete(c.responses, rid)

	ctx, cancel := context.WithTimeout(ctx, clientTimeout)
	defer cancel()

	select {
	case res := <-c.responses[rid]:
		return res, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *NickServClient) Command(ctx context.Context, msg *pb.NickServRPCCommand) (*pb.NickServRPCResponse, error) {
	rid, err := c.registerResponseChan()
	if err != nil {
		return nil, err
	}

	msg.RequestId = rid

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	addr, ok := c.hostAddr.Load().(*hostAddr)
	if !ok {
		return nil, errors.New("unable to get host address")
	}

	// TODO: correct?
	err = c.svc.Network.Send(addr.HostID, addr.Port, c.port, msgBytes)
	if err != nil {
		return nil, err
	}

	return c.awaitResponse(ctx, rid)
}

func newNickServItem(r *pb.NickServRecord) *nickServItem {
	return &nickServItem{
		key:    r.Key,
		record: unsafe.Pointer(r),
	}
}

type nickServItem struct {
	key    []byte
	record unsafe.Pointer
}

func (p *nickServItem) Record() *pb.NickServRecord {
	return (*pb.NickServRecord)(atomic.LoadPointer(&p.record))
}

// Less implements llrb.Item
func (p *nickServItem) Less(oi llrb.Item) bool {

	if o, ok := oi.(*nickServItem); ok {
		return bytes.Compare(p.key, o.key) == -1
	}
	return !oi.Less(p)
}

func (s *NickServStore) Insert(r *pb.NickServRecord) error {
	s.Lock()
	defer s.Unlock()

	item := newNickServItem(r)

	_, ok := s.records.Get(item).(*nickServItem)
	if ok {
		return ErrAlreadyExists
	}

	s.records.ReplaceOrInsert(item)
	s.nicks[strings.ToLower(r.Nick)] = item
	return nil
}

func (s *NickServStore) Delete(key []byte) error {
	s.Lock()
	defer s.Unlock()

	keyItem := &nickServItem{key: key}

	if s.records.Get(keyItem) != nil {
		s.records.Delete(keyItem)
		return nil
	}

	return ErrNotFound
}

func (s *NickServStore) Update(r *pb.NickServRecord, oldNick string) error {
	s.Lock()
	defer s.Unlock()

	prev, ok := s.records.Get(&nickServItem{key: r.Key}).(*nickServItem)
	if !ok {
		return ErrNotFound
	}

	// update nick index if required
	if oldNick != r.Nick {
		delete(s.nicks, strings.ToLower(oldNick))
		s.nicks[strings.ToLower(r.Nick)] = prev
	}

	prev.SetRecord(r)

	return nil
}

func (p *nickServItem) SetRecord(r *pb.NickServRecord) {
	atomic.SwapPointer(&p.record, unsafe.Pointer(r))
}

func (s *NickServStore) Retrieve(key []byte) (*pb.NickServRecord, error) {
	s.Lock()
	defer s.Unlock()
	item := s.records.Get(&nickServItem{key: key})

	if item == nil {
		return nil, ErrNotFound
	}
	return item.(*nickServItem).Record(), nil
}

func signNickToken(token *pb.NickServToken, key *pb.Key) error {
	tokenBytes, _ := serializeNickToken(token)

	switch key.Type {
	case pb.KeyType_KEY_TYPE_ED25519:
		if len(key.Private) != ed25519.PrivateKeySize {
			return dao.ErrInvalidKeyLength
		}
		token.Signature = ed25519.Sign(key.Private, tokenBytes)
		return nil
	default:
		return dao.ErrUnsupportedKeyType
	}
}

func verifyNickServToken(token *pb.NickServToken) (bool, error) {
	tokenBytes, _ := serializeNickToken(token)
	if len(token.Key) != ed25519.PublicKeySize {
		return false, dao.ErrInvalidKeyLength
	}
	return ed25519.Verify(token.Key, tokenBytes, token.Signature), nil
}

var nextSnowflakeID uint64

func (s *NickServ) validUntil() uint64 {
	return uint64(time.Now().UTC().Unix()) + s.validUntil()
}
