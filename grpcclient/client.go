package grpcclient

import (
	"context"
	"fmt"
	"net"

	"github.com/gotti/meshover/spec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
  OverlayIP net.IP
	conn     *grpc.ClientConn
	GrpcConn spec.ControlPlaneServiceClient
}

func NewClient(hostName string, controlserver string, publickey []byte, underlayAddress string) (*Client, error) {
	conn, err := grpc.Dial(controlserver, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to controlserver, err=%w", err)
	}
	c := spec.NewControlPlaneServiceClient(conn)
  ctx := context.Background()
  adrreq := new(spec.AddressAssignRequest)
  adrreq.Name = hostName
	r, err := c.AddressAssign(ctx, adrreq)
	if err != nil {
		return nil, fmt.Errorf("failed to request address assign, err=%w", err)
	}
  overlayIP := net.ParseIP(r.GetAddress().GetEndpointAddress())
	fmt.Printf("assigned address %s\n", overlayIP)
	req := new(spec.RegisterPeerRequest)
	req.Peer = new(spec.Peer)
	req.Peer.Address = r.GetAddress()
	req.Peer.Asn = "8888"
	req.Peer.Name = hostName
	wg := new(spec.Peer_UnderlayWireguard)
	wg.UnderlayWireguard = new(spec.UnderlayWireguard)
	wg.UnderlayWireguard.Endpoint = new(spec.Address)
	wg.UnderlayWireguard.Endpoint.EndpointAddress = fmt.Sprintf("[%s]:12912", underlayAddress)
	wg.UnderlayWireguard.PublicKey = publickey
	req.Peer.Underlay = wg
	res, err := c.RegisterPeer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to register peer, err=%w", err)
	}
	if !res.Ok {
		return nil, fmt.Errorf("ok is false when registering peer")
	}
  return &Client{OverlayIP: overlayIP, conn:conn, GrpcConn:  c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
