package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"

	"github.com/gotti/meshover/spec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var (
	listenAddress      = flag.String("listen", "", "example: 0.0.0.0:8080")
	stateFilePath      = flag.String("statefile", "state", "filename")
	assignAddressRange = flag.String("assignaddress", "10.128.0.0/12", "cidr ipaddress")
)

type controlServer struct {
	stateFilePath string
	mtx           sync.Mutex
	spec.UnimplementedControlPlaneServiceServer
	peers *spec.Peers
}

func (c *controlServer) SaveState() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	buf, err := proto.Marshal(c.peers)
	if err != nil {
		return fmt.Errorf("failed to marshal for saving state, err=%w", err)
	}
	if err := os.WriteFile(c.stateFilePath, buf, 0600); err != nil {
		return fmt.Errorf("failed to write state, err=%w", err)
	}
	return nil
}

func (c *controlServer) LoadState() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, err := os.Stat(c.stateFilePath); os.IsNotExist(err) {
		f, err := os.Create(c.stateFilePath)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("failed to create state file")
		}
	}
	f, err := os.ReadFile(c.stateFilePath)
	if err != nil {
		return fmt.Errorf("failed to load state, err=%w", err)
	}
	buf := new(spec.Peers)
	if err := proto.Unmarshal(f, buf); err != nil {
		return fmt.Errorf("failed to unmarshal for loading state, err=%w", err)
	}
	c.peers = buf
	return nil
}

func (c *controlServer) AddPeer(pe *spec.Peer) {
	for i, p := range c.peers.Peers {
		if p.GetName() == pe.GetName() {
			c.peers.Peers[i] = pe
			return
		}
	}
	c.peers.Peers = append(c.peers.Peers, pe)
}

func (c *controlServer) FindPeer(name string) *spec.Peer {
	for _, p := range c.peers.Peers {
		if p.GetName() == name {
			return p
		}
	}
	return nil
}

func (c *controlServer) ListPeers(ctx context.Context, in *spec.ListPeersRequest) (*spec.ListPeersResponse, error) {
	ret := new(spec.ListPeersResponse)
	ret.Peers = c.peers
	c.SaveState()
	return ret, nil
}

func (c *controlServer) AddressAssign(ctx context.Context, in *spec.AddressAssignRequest) (*spec.AddressAssignResponse, error) {
	p := c.FindPeer(in.GetName())
	if p != nil {
		ret := &spec.AddressAssignResponse{Address: p.GetAddress(), Asnumber: p.Asnumber}
		return ret, nil
	}
	randomAddr := uint32(rand.Int())
	i, n, err := net.ParseCIDR(*assignAddressRange)
	if err != nil {
		return nil, err
	}
	m, b := n.Mask.Size()
	if b != 32 {
		log.Fatalln("unknown size")
	}
	addr := (randomAddr & ((1 << (32 - m)) - 1)) + binary.BigEndian.Uint32(i.Mask(n.Mask))
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, addr)
	fmt.Printf("assigned %s\n", ip.String())
	randomASN := rand.Intn(94967294) + 4200000000
	ret := spec.AddressAssignResponse{Address: spec.NewAddress(ip), Asnumber: &spec.ASN{Number: uint32(randomASN)}}
	return &ret, nil
}

func (c *controlServer) RegisterPeer(ctx context.Context, in *spec.RegisterPeerRequest) (*spec.RegisterPeerResponse, error) {
	p := in.GetPeer()
	c.AddPeer(p)
	fmt.Printf("%+v", c.peers)
	if err := c.SaveState(); err != nil {
		log.Fatalf("failed to save, %s", err)
	}
	ret := new(spec.RegisterPeerResponse)
	ret.Ok = true
	return ret, nil
}

func loadAndSanitizeArgs() {
	flag.Parse()
	if _, _, err := net.ParseCIDR(*assignAddressRange); err != nil {
		log.Fatalf("unrecognized assignAddressRange, err=%s", err)
	}
}

type authorizer struct {
	token  string
	server *controlServer
}

func (a *authorizer) Context(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	if info.FullMethod == "/ControlPlaneService/AddressAssign" {
	}
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: no context metadata")
	}
	return ctx, nil
}

func (a *authorizer) HandleUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := a.Context(ctx, info)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func main() {
	loadAndSanitizeArgs()
	server := controlServer{}
	server.stateFilePath = *stateFilePath
	if err := server.LoadState(); err != nil {
		log.Fatalf("failed to load state, err=%s\n", err)
	}
	fmt.Printf("%+v\n", server.peers)
	//a := authorizer{token: "a", server: &server}
	s := grpc.NewServer(
	//grpc.UnaryInterceptor(a.HandleUnary),
	)
	spec.RegisterControlPlaneServiceServer(s, &server)
	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("failed to listen, err=%s\n", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve, err=%s\n", err)
	}
}
