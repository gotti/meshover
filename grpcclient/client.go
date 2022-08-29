package grpcclient

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gotti/meshover/spec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	context   context.Context
	OverlayIP net.IPNet
	conn      *grpc.ClientConn
	grpcConn  spec.ControlPlaneServiceClient
}

func NewClient(ctx context.Context, hostName string, controlserver string, publickey *spec.Curve25519Key, underlayAddress *spec.AddressAndPort) (*Client, *spec.ASN, error) {
	conn, err := grpc.Dial(controlserver, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to controlserver, err=%w", err)
	}
	c := spec.NewControlPlaneServiceClient(conn)
	adrreq := new(spec.AddressAssignRequest)
	adrreq.Name = hostName
	r, err := c.AddressAssign(ctx, adrreq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to request address assign, err=%w", err)
	}
	overlayIP := r.GetAddress()[0]
	overlayIPNet := overlayIP.ToNetIPNet()
	if overlayIPNet == nil {
		log.Fatalln("failed to parse overlayIP")
	}
	fmt.Printf("assigned address %s\n", r.GetAddress())
	req := new(spec.RegisterPeerRequest)
	req.Peer = new(spec.Peer)
	req.Peer.Address = r.GetAddress()
	req.Peer.Asnumber = r.GetAsnumber()
	req.Peer.Name = hostName
	wg := &spec.UnderlayLinuxKernelWireguard{Endpoint: underlayAddress, PublicKey: publickey}
	req.Peer.Underlay = &spec.Peer_UnderlayLinuxKernelWireguard{UnderlayLinuxKernelWireguard: wg}
	res, err := c.RegisterPeer(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to register peer, err=%w", err)
	}
	if !res.Ok {
		return nil, nil, fmt.Errorf("ok is false when registering peer")
	}
	return &Client{context: ctx, OverlayIP: *overlayIPNet, conn: conn, grpcConn: c}, r.GetAsnumber(), nil
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
