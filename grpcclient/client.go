package grpcclient

import (
	"context"
	"fmt"
	"net"

	"github.com/gotti/meshover/spec"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	context       context.Context
	OverlayIP     *net.IPNet
	AdditionalIPs []*net.IPNet
	conn          *grpc.ClientConn
	grpcConn      spec.ControlPlaneServiceClient
}

func NewClient(ctx context.Context, log *zap.Logger, hostName string, controlserver string, agentToken string, publickey *spec.Curve25519Key, underlayAddress *spec.AddressAndPort, routeGather []*net.IPNet) (*Client, *spec.ASN, error) {
	conn, err := grpc.Dial(controlserver, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to controlserver, err=%w", err)
	}
	c := spec.NewControlPlaneServiceClient(conn)
	adrreq := new(spec.AddressAssignRequest)
	adrreq.Name = hostName
	md := metadata.New(map[string]string{"Bearer": agentToken})
	ctx = metadata.NewOutgoingContext(ctx, md)
	r, err := c.AddressAssign(ctx, adrreq)
	fmt.Printf("assigned address %s, wg address %s\n", r.GetAddress(), r.GetWireguardAddress())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to request address assign, err=%w", err)
	}
	overlayIP := r.GetWireguardAddress().ToNetIPNet()
	var additionalIPs []*net.IPNet
	for i := range r.GetAddress() {
		additionalIPs = append(additionalIPs, r.GetAddress()[i].ToNetIPNet())
	}
	sbr := &spec.SourceBasedRoutingOption{SourceIPRange: []*spec.AddressCIDR{}}
	for _, r := range routeGather {
		sbr.SourceIPRange = append(sbr.SourceIPRange, spec.NewAddressCIDR(*r))
	}
	req := &spec.RegisterPeerRequest{
		Peer: &spec.Peer{
			WireguardAddress: r.GetWireguardAddress(),
			Underlay: &spec.Peer_UnderlayLinuxKernelWireguard{
				UnderlayLinuxKernelWireguard: &spec.UnderlayLinuxKernelWireguard{
					Endpoint:  underlayAddress,
					PublicKey: publickey,
				},
			},
			Address:   r.GetAddress(),
			Asnumber:  r.GetAsnumber(),
			Name:      hostName,
			SbrOption: sbr,
		},
	}
	res, err := c.RegisterPeer(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to register peer, err=%w", err)
	}
	if !res.Ok {
		return nil, nil, fmt.Errorf("ok is false when registering peer")
	}
	return &Client{context: ctx, OverlayIP: overlayIP, AdditionalIPs: additionalIPs, conn: conn, grpcConn: c}, r.GetAsnumber(), nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ListPeers() (*spec.ListPeersResponse, error) {
	req := new(spec.ListPeersRequest)
	ret, err := c.grpcConn.ListPeers(c.context, req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect ListPeers endpont, err=%w", err)
	}
	if err := ret.ValidateAll(); err != nil {
		return nil, fmt.Errorf("validation failed, err=%w", err)
	}
	return ret, nil
}
