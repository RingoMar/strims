// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package dao

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/MemeLabs/protobuf/pkg/rpc"
	"github.com/MemeLabs/strims/internal/dao/versionvector"
	networkv1 "github.com/MemeLabs/strims/pkg/apis/network/v1"
	networkv1bootstrap "github.com/MemeLabs/strims/pkg/apis/network/v1/bootstrap"
	networkv1ca "github.com/MemeLabs/strims/pkg/apis/network/v1/ca"
	networkv1directory "github.com/MemeLabs/strims/pkg/apis/network/v1/directory"
	networkv1errors "github.com/MemeLabs/strims/pkg/apis/network/v1/errors"
	profilev1 "github.com/MemeLabs/strims/pkg/apis/profile/v1"
	"github.com/MemeLabs/strims/pkg/apis/type/certificate"
	"github.com/MemeLabs/strims/pkg/apis/type/image"
	"github.com/MemeLabs/strims/pkg/apis/type/key"
	"github.com/MemeLabs/strims/pkg/hashmap"
	"github.com/MemeLabs/strims/pkg/kv"
	"github.com/MemeLabs/strims/pkg/timeutil"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	_ = iota + networkNS
	networkNetworkNS
	networkNetworkKeyNS
	networkCertificateLogNS
	networkCertificateLogNetworkNS
	networkCertificateLogSerialNS
	networkCertificateLogSubjectNS
	networkBootstrapClientNS
	networkUIConfigNS
	networkPeerNS
	networkPeerNetworkNS
	networkPeerPublicKeyNS
	networkPeerInviterNS
	networkCertificateLogKeyNS
	networkBootstrapClientClientOptionsNS
	bootstrapConfigNS
	networkAliasReservationNS
	networkAliasReservationsByNetworkIDNS
	networkAliasReservationsByAliasNS
)

// TODO: move to network server config
const (
	NetworkAliasChangeCooldown      = 30 * 24 * time.Hour
	NetworkAliasReservationCooldown = 180 * 24 * time.Hour
)

var Networks = NewTable(
	networkNetworkNS,
	&TableOptions[networkv1.Network, *networkv1.Network]{
		ObserveChange: func(m, p *networkv1.Network) proto.Message {
			return &networkv1.NetworkChangeEvent{Network: m}
		},
		ObserveDelete: func(m *networkv1.Network) proto.Message {
			return &networkv1.NetworkDeleteEvent{Network: m}
		},
		OnChange: LocalSetHook(func(s ReplicatedRWTx, m, p *networkv1.Network) error {
			if p != nil || m.GetServerConfig() == nil {
				return nil
			}

			for c := m.Certificate; c != nil; c = c.GetParent() {
				log, err := NewCertificateLog(ProfileID.IDGenerator(s), m.Id, c)
				if err != nil {
					return err
				}
				if err := CertificateLogs.Insert(s, log); err != nil {
					return err
				}
			}

			directoryUser, err := NewDirectoryUserRecord(ProfileID.IDGenerator(s), m.Id, m.Certificate.Key)
			if err != nil {
				return err
			}
			directoryUser.Moderation.IsAdmin = wrapperspb.Bool(true)
			directoryUser.Moderation.IsModerator = wrapperspb.Bool(true)
			return DirectoryUserRecords.Insert(s, directoryUser)
		}),
	},
)

func resolveNetworkConflict(m, p *networkv1.Network) {
	if m.Version.UpdatedAt < p.Version.UpdatedAt {
		m.Alias = p.Alias
		m.CertificateRenewalError = p.CertificateRenewalError
	}
	if m.Certificate.NotAfter < p.Certificate.NotAfter {
		m.Certificate = p.Certificate
	}
	m.ServerConfig = p.ServerConfig
	versionvector.Upgrade(m.GetVersion(), p.GetVersion())
}

func init() {
	RegisterReplicatedTable(
		Networks,
		&ReplicatorOptions[*networkv1.Network]{
			OnConflict: func(s kv.RWStore, m *networkv1.Network, p *networkv1.Network) error {
				resolveNetworkConflict(m, p)
				return Networks.Update(s, m)
			},
			Extract: func(s kv.Store, m, p *networkv1.Network) *networkv1.Network {
				m.ServerConfig = nil
				return m
			},
			Merge: func(s kv.RWStore, m, p *networkv1.Network) *networkv1.Network {
				versionvector.Upgrade(m.GetVersion(), p.GetVersion())
				m.ServerConfig = p.ServerConfig
				return m
			},
		},
	)
}

var NetworksByKey = NewUniqueIndex(
	networkNetworkKeyNS,
	Networks,
	NetworkKey,
	byteIdentity,
	&UniqueIndexOptions[networkv1.Network, *networkv1.Network]{
		OnConflict: func(s kv.RWStore, t *Table[networkv1.Network, *networkv1.Network], m, p *networkv1.Network) error {
			// TODO: we don't currently have good support for replicating records that
			// have both 1:many edges and unique keys...
			//
			// the only case where this would happen is if replia A creates a network
			// then replica B accepts an invite to the same network before reading A's
			// replication log.
			if p.ServerConfig != nil {
				return ErrUniqueConstraintViolated
			}
			resolveNetworkConflict(m, p)
			return Networks.Delete(s, p.Id)
		},
	},
)

// NewNetworkCertificate ...
func NewNetworkCertificate(config *networkv1.ServerConfig) (*certificate.Certificate, error) {
	return NewSelfSignedCertificate(config.Key, certificate.KeyUsage_KEY_USAGE_SIGN, defaultCertTTL, WithSubject(config.Name))
}

func SignCertificateRequestWithNetwork(csr *certificate.CertificateRequest, config *networkv1.ServerConfig) (*certificate.Certificate, error) {
	networkCert, err := NewNetworkCertificate(config)
	if err != nil {
		return nil, err
	}

	cert, err := SignCertificateRequest(csr, defaultCertTTL, config.Key)
	if err != nil {
		return nil, err
	}
	cert.ParentOneof = &certificate.Certificate_Parent{Parent: networkCert}

	return cert, nil
}

type NewNetworkOptions struct {
	CertificateRequestOptions []CertificateRequestOption
	Alias                     string
}

type NewNetworkOption func(o *NewNetworkOptions)

func WithCertificateRequestOption(opt CertificateRequestOption) NewNetworkOption {
	return func(o *NewNetworkOptions) {
		o.CertificateRequestOptions = append(o.CertificateRequestOptions, opt)
	}
}

func WithAlias(alias string) NewNetworkOption {
	return func(o *NewNetworkOptions) {
		if alias != "" {
			o.Alias = alias
			o.CertificateRequestOptions = append(o.CertificateRequestOptions, WithSubject(alias))
		}
	}
}

// NewNetwork ...
func NewNetwork(g IDGenerator, name string, icon *image.Image, profile *profilev1.Profile, opts ...NewNetworkOption) (*networkv1.Network, error) {
	o := &NewNetworkOptions{
		CertificateRequestOptions: []CertificateRequestOption{
			WithSubject(profile.Name),
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	key, err := GenerateKey()
	if err != nil {
		return nil, err
	}

	network := &networkv1.Network{
		Id:      id,
		Version: versionvector.New(),
		ServerConfig: &networkv1.ServerConfig{
			Name: name,
			Key:  key,
			Directory: &networkv1directory.ServerConfig{
				Integrations: &networkv1directory.ServerConfig_Integrations{
					Swarm: &networkv1directory.ServerConfig_Integrations_Swarm{
						Enable: true,
					},
				},
			},
			Icon: icon,
		},
	}

	csr, err := NewCertificateRequest(
		profile.Key,
		certificate.KeyUsage_KEY_USAGE_PEER|certificate.KeyUsage_KEY_USAGE_SIGN,
		o.CertificateRequestOptions...,
	)
	if err != nil {
		return nil, err
	}

	cert, err := SignCertificateRequestWithNetwork(csr, network.GetServerConfig())
	if err != nil {
		return nil, err
	}
	network.Certificate = cert

	if profile.Name != cert.Subject {
		network.Alias = cert.Subject
	}

	return network, nil
}

// NewNetworkFromInvitationV0 generates a network from a network invitation
func NewNetworkFromInvitationV0(g IDGenerator, invitation *networkv1.InvitationV0, profile *profilev1.Profile, opts ...NewNetworkOption) (*networkv1.Network, error) {
	o := &NewNetworkOptions{
		CertificateRequestOptions: []CertificateRequestOption{
			WithSubject(profile.Name),
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	if err = VerifyCertificate(invitation.Certificate); err != nil {
		return nil, err
	}
	csr, err := NewCertificateRequest(profile.Key, certificate.KeyUsage_KEY_USAGE_PEER, o.CertificateRequestOptions...)
	if err != nil {
		return nil, err
	}
	peerCert, err := SignCertificateRequest(csr, defaultCertTTL, invitation.Key)
	if err != nil {
		return nil, err
	}
	peerCert.ParentOneof = &certificate.Certificate_Parent{Parent: invitation.Certificate}

	return &networkv1.Network{
		Id:          id,
		Version:     versionvector.New(),
		Certificate: peerCert,
		Alias:       o.Alias,
	}, nil
}

// NewNetworkFromCertificate generates a network from a network invitation
func NewNetworkFromCertificate(g IDGenerator, cert *certificate.Certificate) (*networkv1.Network, error) {
	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	if err = VerifyCertificate(cert); err != nil {
		return nil, err
	}

	return &networkv1.Network{
		Id:          id,
		Version:     versionvector.New(),
		Certificate: cert,
	}, nil
}

// NewInvitationV0 ...
func NewInvitationV0(key *key.Key, cert *certificate.Certificate, bootstrapClients []*networkv1bootstrap.BootstrapClient) (*networkv1.InvitationV0, error) {
	inviteKey, err := GenerateKey()
	if err != nil {
		return nil, err
	}

	validDuration := time.Hour * 24 * 7 // seven days

	inviteCSR, err := NewCertificateRequest(inviteKey, certificate.KeyUsage_KEY_USAGE_SIGN)
	if err != nil {
		return nil, err
	}
	inviteCert, err := SignCertificateRequest(inviteCSR, validDuration, key)
	if err != nil {
		return nil, err
	}
	inviteCert.ParentOneof = &certificate.Certificate_Parent{Parent: cert}

	invitation := &networkv1.InvitationV0{
		Key:         inviteKey,
		Certificate: inviteCert,
	}

	for _, c := range bootstrapClients {
		cc := proto.Clone(c).(*networkv1bootstrap.BootstrapClient)
		cc.Id = 0
		invitation.BootstrapClients = append(invitation.BootstrapClients, cc)
	}

	return invitation, nil
}

func CertificateChain(cert *certificate.Certificate) []*certificate.Certificate {
	var chain []*certificate.Certificate
	for ; cert != nil; cert = cert.GetParent() {
		chain = append(chain, cert)
	}
	return chain
}

var CertificateLogs = NewTable[networkv1ca.CertificateLog](networkCertificateLogNS, nil)

var CertificateLogsByNetwork = ManyToOne(
	networkCertificateLogNetworkNS,
	CertificateLogs,
	Networks,
	(*networkv1ca.CertificateLog).GetNetworkId,
	&ManyToOneOptions[networkv1ca.CertificateLog, *networkv1ca.CertificateLog]{CascadeDelete: true},
)

func FormatCertificateLogSerialNumberKey(networkID uint64, serialNumber []byte) []byte {
	b := make([]byte, 8, 8+len(serialNumber))
	binary.BigEndian.PutUint64(b, networkID)
	return append(b, serialNumber...)
}

func certificateLogSerialNumberKey(m *networkv1ca.CertificateLog) []byte {
	return FormatCertificateLogSerialNumberKey(m.NetworkId, m.Certificate.SerialNumber)
}

var CertificateLogsBySerialNumber = NewUniqueIndex(
	networkCertificateLogSerialNS,
	CertificateLogs,
	certificateLogSerialNumberKey,
	byteIdentity,
	nil,
)

func FormatCertificateLogSubjectKey(networkID uint64, subject string) []byte {
	b := make([]byte, 8, 8+len([]byte(subject)))
	binary.BigEndian.PutUint64(b, networkID)
	return append(b, []byte(subject)...)
}

func certificateLogSubjectKey(m *networkv1ca.CertificateLog) []byte {
	return FormatCertificateLogSubjectKey(m.NetworkId, m.Certificate.Subject)
}

var ErrCertificateSubjectInUse = errors.New("certificate subject in use")

var CertificateLogsBySubject = NewUniqueIndex(
	networkCertificateLogSubjectNS,
	CertificateLogs,
	certificateLogSubjectKey,
	byteIdentity,
	&UniqueIndexOptions[networkv1ca.CertificateLog, *networkv1ca.CertificateLog]{
		OnConflict: func(s kv.RWStore, t *Table[networkv1ca.CertificateLog, *networkv1ca.CertificateLog], m, p *networkv1ca.CertificateLog) error {
			return DeleteSecondaryIndex(s, networkCertificateLogSubjectNS, certificateLogSubjectKey(m), p.Id)
		},
	},
)

func FormatCertificateLogKeyKey(networkID uint64, key []byte) []byte {
	b := make([]byte, 8, 8+len([]byte(key)))
	binary.BigEndian.PutUint64(b, networkID)
	return append(b, []byte(key)...)
}

func certificateLogKeyKey(m *networkv1ca.CertificateLog) []byte {
	return FormatCertificateLogKeyKey(m.NetworkId, m.Certificate.Key)
}

var CertificateLogsByKey = NewUniqueIndex(
	networkCertificateLogKeyNS,
	CertificateLogs,
	certificateLogKeyKey,
	byteIdentity,
	&UniqueIndexOptions[networkv1ca.CertificateLog, *networkv1ca.CertificateLog]{
		OnConflict: func(s kv.RWStore, t *Table[networkv1ca.CertificateLog, *networkv1ca.CertificateLog], m, p *networkv1ca.CertificateLog) error {
			return DeleteSecondaryIndex(s, networkCertificateLogKeyNS, certificateLogKeyKey(m), p.Id)
		},
	},
)

// NewCertificateLog ...
func NewCertificateLog(s IDGenerator, networkID uint64, cert *certificate.Certificate) (*networkv1ca.CertificateLog, error) {
	id, err := s.GenerateID()
	if err != nil {
		return nil, err
	}

	c := proto.Clone(cert).(*certificate.Certificate)
	if p := c.GetParent(); p != nil {
		c.ParentOneof = &certificate.Certificate_ParentSerialNumber{
			ParentSerialNumber: p.SerialNumber,
		}
	}

	return &networkv1ca.CertificateLog{
		Id:          id,
		NetworkId:   networkID,
		Certificate: c,
	}, nil
}

func NewCertificateLogCache(s kv.RWStore, opt *CacheStoreOptions) (c CertificateLogCache) {
	c.CacheStore, c.ByID = newCacheStore[networkv1ca.CertificateLog](s, CertificateLogs, opt)
	c.BySerialNumber = NewCacheIndex(
		c.CacheStore,
		CertificateLogsBySerialNumber.Get,
		certificateLogSerialNumberKey,
		hashmap.NewByteInterface[[]byte],
	)
	c.BySubject = NewCacheIndex(
		c.CacheStore,
		CertificateLogsBySubject.Get,
		certificateLogSubjectKey,
		hashmap.NewByteInterface[[]byte],
	)
	c.ByKey = NewCacheIndex(
		c.CacheStore,
		CertificateLogsByKey.Get,
		certificateLogKeyKey,
		hashmap.NewByteInterface[[]byte],
	)
	return
}

type CertificateLogCache struct {
	*CacheStore[networkv1ca.CertificateLog, *networkv1ca.CertificateLog]
	ByID           CacheAccessor[uint64, networkv1ca.CertificateLog, *networkv1ca.CertificateLog]
	BySerialNumber CacheAccessor[[]byte, networkv1ca.CertificateLog, *networkv1ca.CertificateLog]
	BySubject      CacheAccessor[[]byte, networkv1ca.CertificateLog, *networkv1ca.CertificateLog]
	ByKey          CacheAccessor[[]byte, networkv1ca.CertificateLog, *networkv1ca.CertificateLog]
}

// NetworkKey ...
func NetworkKey(network *networkv1.Network) []byte {
	return CertificateRoot(network.Certificate).Key
}

var BootstrapConfig = NewSingleton(
	bootstrapConfigNS,
	&SingletonOptions[networkv1bootstrap.Config, *networkv1bootstrap.Config]{
		DefaultValue: &networkv1bootstrap.Config{
			EnablePublishing: false,
		},
		ObserveChange: func(m, p *networkv1bootstrap.Config) proto.Message {
			return &networkv1bootstrap.ConfigChangeEvent{Config: m}
		},
	},
)

var BootstrapClients = NewTable(
	networkBootstrapClientNS,
	&TableOptions[networkv1bootstrap.BootstrapClient, *networkv1bootstrap.BootstrapClient]{
		ObserveChange: func(m, p *networkv1bootstrap.BootstrapClient) proto.Message {
			return &networkv1bootstrap.BootstrapClientChange{BootstrapClient: m}
		},
	},
)

func init() {
	RegisterReplicatedTable(BootstrapClients, nil)
}

func FormatBootstrapClientClientOptionsKey(m *networkv1bootstrap.BootstrapClient) []byte {
	switch o := m.ClientOptions.(type) {
	case *networkv1bootstrap.BootstrapClient_WebsocketOptions:
		return []byte(o.WebsocketOptions.Url)
	default:
		return nil
	}
}

var BootstrapClientsByClientOptions = NewUniqueIndex(
	networkBootstrapClientClientOptionsNS,
	BootstrapClients,
	FormatBootstrapClientClientOptionsKey,
	byteIdentity,
	&UniqueIndexOptions[networkv1bootstrap.BootstrapClient, *networkv1bootstrap.BootstrapClient]{
		OnConflict: func(s kv.RWStore, t *Table[networkv1bootstrap.BootstrapClient, *networkv1bootstrap.BootstrapClient], m, p *networkv1bootstrap.BootstrapClient) error {
			return BootstrapClients.Delete(s, p.GetId())
		},
	},
)

// NewWebSocketBootstrapClient ...
func NewWebSocketBootstrapClient(g IDGenerator, url string, insecureSkipVerifyTLS bool) (*networkv1bootstrap.BootstrapClient, error) {
	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	return &networkv1bootstrap.BootstrapClient{
		Id: id,
		ClientOptions: &networkv1bootstrap.BootstrapClient_WebsocketOptions{
			WebsocketOptions: &networkv1bootstrap.BootstrapClientWebSocketOptions{
				Url:                   url,
				InsecureSkipVerifyTls: insecureSkipVerifyTLS,
			},
		},
	}, nil
}

// NewBootstrapClient ...
func NewBootstrapClient(g IDGenerator, bootstrapClient *networkv1bootstrap.BootstrapClient) (*networkv1bootstrap.BootstrapClient, error) {
	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	return &networkv1bootstrap.BootstrapClient{
		Id:            id,
		ClientOptions: bootstrapClient.ClientOptions,
	}, nil
}

var NetworkUIConfig = NewSingleton(
	networkUIConfigNS,
	&SingletonOptions[networkv1.UIConfig, *networkv1.UIConfig]{
		DefaultValue: &networkv1.UIConfig{
			NetworkDisplayOrder: []uint64{},
		},
		ObserveChange: func(m, p *networkv1.UIConfig) proto.Message {
			return &networkv1.UIConfigChangeEvent{UiConfig: m}
		},
	},
)

var NetworkPeers = NewTable[networkv1.Peer](networkPeerNS, nil)

var NetworkPeersByNetwork = ManyToOne(
	networkPeerNetworkNS,
	NetworkPeers,
	Networks,
	(*networkv1.Peer).GetNetworkId,
	&ManyToOneOptions[networkv1.Peer, *networkv1.Peer]{CascadeDelete: true},
)

var NetworkPeersByInviterPeer = ManyToOne(
	networkPeerInviterNS,
	NetworkPeers,
	NetworkPeers,
	(*networkv1.Peer).GetInviterPeerId,
	nil,
)

func FormatNetworkPeerPublicKeyKey(networkID uint64, key []byte) []byte {
	b := make([]byte, 8, 8+len(key))
	binary.BigEndian.PutUint64(b, networkID)
	return append(b, key...)
}

var NetworkPeersByPublicKey = NewUniqueIndex(
	networkPeerPublicKeyNS,
	NetworkPeers,
	func(m *networkv1.Peer) []byte {
		return FormatNetworkPeerPublicKeyKey(m.NetworkId, m.PublicKey)
	},
	byteIdentity,
	nil,
)

// NewNetworkPeer ...
func NewNetworkPeer(g IDGenerator, networkID uint64, publicKey []byte, alias string, inviterPeerID uint64) (*networkv1.Peer, error) {
	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	return &networkv1.Peer{
		Id:            id,
		NetworkId:     networkID,
		PublicKey:     publicKey,
		Alias:         alias,
		InviterPeerId: inviterPeerID,
		InviteQuota:   3,
		CreatedAt:     timeutil.Now().Unix(),
	}, nil
}

type NetworkAliasReservationTable struct {
	*Table[networkv1.AliasReservation, *networkv1.AliasReservation]
}

func (t *NetworkAliasReservationTable) Reserve(s kv.RWStore, id uint64, peerKey []byte) error {
	_, err := t.Table.Transform(s, id, func(p *networkv1.AliasReservation) error {
		if bytes.Equal(p.PeerKey, peerKey) {
			return nil
		}
		if !timeutil.Now().After(timeutil.Unix(p.ReservedUntil, 0)) {
			return rpc.WrapError(errors.New("alias in use"), networkv1errors.ErrorCode_ALIAS_IN_USE)
		}
		p.PeerKey = peerKey
		p.ReservedUntil = timeutil.MaxTime.Unix()
		return nil
	})
	return err
}

func (t *NetworkAliasReservationTable) Release(s kv.RWStore, id uint64, cooldown time.Duration) error {
	_, err := t.Table.Transform(s, id, func(p *networkv1.AliasReservation) error {
		p.PeerKey = nil
		p.ReservedUntil = timeutil.Now().Add(cooldown).Unix()
		return nil
	})
	return err
}

var NetworkAliasReservations = NetworkAliasReservationTable{NewTable[networkv1.AliasReservation](networkAliasReservationNS, nil)}

var NetworkAliasReservationsByNetwork = ManyToOne(
	networkAliasReservationsByNetworkIDNS,
	NetworkAliasReservations.Table,
	Networks,
	(*networkv1.AliasReservation).GetNetworkId,
	nil,
)

func FormatNetworkAliasReservationAliasKey(networkID uint64, alias string) []byte {
	b := make([]byte, 8, 8+len(alias))
	binary.BigEndian.PutUint64(b, networkID)
	return append(b, alias...)
}

var NetworkAliasReservationsByAlias = NewUniqueIndex(
	networkAliasReservationsByAliasNS,
	NetworkAliasReservations.Table,
	func(m *networkv1.AliasReservation) []byte {
		return FormatNetworkAliasReservationAliasKey(m.NetworkId, m.Alias)
	},
	byteIdentity,
	nil,
)

func NewNetworkAliasReservation(g IDGenerator, networkID uint64, alias string, peerKey []byte) (*networkv1.AliasReservation, error) {
	id, err := g.GenerateID()
	if err != nil {
		return nil, err
	}

	return &networkv1.AliasReservation{
		Id:            id,
		NetworkId:     networkID,
		Alias:         alias,
		PeerKey:       peerKey,
		ReservedUntil: timeutil.MaxTime.Unix(),
	}, nil
}
