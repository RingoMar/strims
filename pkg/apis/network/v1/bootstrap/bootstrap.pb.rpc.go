package bootstrap

import (
	"context"

	"github.com/MemeLabs/protobuf/pkg/rpc"
)

// RegisterBootstrapFrontendService ...
func RegisterBootstrapFrontendService(host rpc.ServiceRegistry, service BootstrapFrontendService) {
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.CreateClient", service.CreateClient)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.UpdateClient", service.UpdateClient)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.DeleteClient", service.DeleteClient)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.GetClient", service.GetClient)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.ListClients", service.ListClients)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.ListPeers", service.ListPeers)
	host.RegisterMethod("strims.network.v1.bootstrap.BootstrapFrontend.PublishNetworkToPeer", service.PublishNetworkToPeer)
}

// BootstrapFrontendService ...
type BootstrapFrontendService interface {
	CreateClient(
		ctx context.Context,
		req *CreateBootstrapClientRequest,
	) (*CreateBootstrapClientResponse, error)
	UpdateClient(
		ctx context.Context,
		req *UpdateBootstrapClientRequest,
	) (*UpdateBootstrapClientResponse, error)
	DeleteClient(
		ctx context.Context,
		req *DeleteBootstrapClientRequest,
	) (*DeleteBootstrapClientResponse, error)
	GetClient(
		ctx context.Context,
		req *GetBootstrapClientRequest,
	) (*GetBootstrapClientResponse, error)
	ListClients(
		ctx context.Context,
		req *ListBootstrapClientsRequest,
	) (*ListBootstrapClientsResponse, error)
	ListPeers(
		ctx context.Context,
		req *ListBootstrapPeersRequest,
	) (*ListBootstrapPeersResponse, error)
	PublishNetworkToPeer(
		ctx context.Context,
		req *PublishNetworkToBootstrapPeerRequest,
	) (*PublishNetworkToBootstrapPeerResponse, error)
}

// BootstrapFrontendClient ...
type BootstrapFrontendClient struct {
	client rpc.Caller
}

// NewBootstrapFrontendClient ...
func NewBootstrapFrontendClient(client rpc.Caller) *BootstrapFrontendClient {
	return &BootstrapFrontendClient{client}
}

// CreateClient ...
func (c *BootstrapFrontendClient) CreateClient(
	ctx context.Context,
	req *CreateBootstrapClientRequest,
	res *CreateBootstrapClientResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.CreateClient", req, res)
}

// UpdateClient ...
func (c *BootstrapFrontendClient) UpdateClient(
	ctx context.Context,
	req *UpdateBootstrapClientRequest,
	res *UpdateBootstrapClientResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.UpdateClient", req, res)
}

// DeleteClient ...
func (c *BootstrapFrontendClient) DeleteClient(
	ctx context.Context,
	req *DeleteBootstrapClientRequest,
	res *DeleteBootstrapClientResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.DeleteClient", req, res)
}

// GetClient ...
func (c *BootstrapFrontendClient) GetClient(
	ctx context.Context,
	req *GetBootstrapClientRequest,
	res *GetBootstrapClientResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.GetClient", req, res)
}

// ListClients ...
func (c *BootstrapFrontendClient) ListClients(
	ctx context.Context,
	req *ListBootstrapClientsRequest,
	res *ListBootstrapClientsResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.ListClients", req, res)
}

// ListPeers ...
func (c *BootstrapFrontendClient) ListPeers(
	ctx context.Context,
	req *ListBootstrapPeersRequest,
	res *ListBootstrapPeersResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.ListPeers", req, res)
}

// PublishNetworkToPeer ...
func (c *BootstrapFrontendClient) PublishNetworkToPeer(
	ctx context.Context,
	req *PublishNetworkToBootstrapPeerRequest,
	res *PublishNetworkToBootstrapPeerResponse,
) error {
	return c.client.CallUnary(ctx, "strims.network.v1.bootstrap.BootstrapFrontend.PublishNetworkToPeer", req, res)
}