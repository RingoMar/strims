package bootstrap

import (
	"github.com/MemeLabs/go-ppspp/internal/api"
	"github.com/MemeLabs/go-ppspp/pkg/syncutil"
	"github.com/MemeLabs/go-ppspp/pkg/vnic"
)

type Peer any

var _ Peer = (*peer)(nil)

// Peer ...
type peer struct {
	vnicPeer       *vnic.Peer
	client         api.PeerClient
	PublishEnabled syncutil.Bool
}
