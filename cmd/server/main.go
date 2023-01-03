package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gotti/meshover/pkg/ip"
	"github.com/gotti/meshover/pkg/keymanager"
	"github.com/gotti/meshover/spec"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	listenAddress         = flag.String("listen", "", "example: 0.0.0.0:8080")
	stateDir              = flag.String("statefile", "state", "filename")
	wireguardAddressRange = flag.String("wgaddress", "10.192.0.0/12", "wireguard peer address")
	assignAddressRange    = flag.String("assignaddress", "10.128.0.0/12,fd00:beef:beef::/48", "list of cidr ipaddress")
	adminPassword         = flag.String("password", "", "strong password")
)

type controlServer struct {
	stateFilePath string
	mtx           sync.Mutex
	logger        *zap.Logger
	spec.UnimplementedControlPlaneServiceServer
	spec.UnimplementedAdministratorServiceServer
	peers  *spec.Peers
	keyman keymanager.Manager
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
			return fmt.Errorf("failed to create state file, err=%w", err)
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
	k, err := keymanager.NewAPIKeys(filepath.Join(*stateDir, "keys"))
	if err != nil {
		return fmt.Errorf("failed to create api keys, err=%w", err)
	}
	c.keyman = k
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
		fmt.Println("found ")
		ret := &spec.AddressAssignResponse{BaseAddress: p.GetBaseAddress(), TunnelAddress: p.GetTunnelAddress(), Asnumber: p.Asnumber}
		return ret, nil
	}
	wgAddress, err := ip.GenerateRandomIP(*wireguardAddressRange)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	addresses := []*spec.AddressCIDR{}
	for _, a := range strings.Split(*assignAddressRange, ",") {
		i, err := ip.GenerateRandomIP(a)
		if err != nil {
			log.Println("failed to generate ip, err=", err)
			return nil, status.Error(codes.Aborted, "failed to generate ip")
		}
		addresses = append(addresses, spec.NewAddressCIDR(*i))
	}
	ret := spec.AddressAssignResponse{WireguardAddress: spec.NewAddressCIDR(*wgAddress), Address: addresses, Asnumber: &spec.ASN{Number: ip.GenerateRandomASN()}}
	return &ret, nil
}

func (c *controlServer) RegisterPeer(ctx context.Context, in *spec.RegisterPeerRequest) (*spec.RegisterPeerResponse, error) {
	p := in.GetPeer()
	c.AddPeer(p)
	fmt.Printf("%+v\n", c.peers)
	if err := c.SaveState(); err != nil {
		log.Fatalf("failed to save, %s", err)
	}
	ret := new(spec.RegisterPeerResponse)
	ret.Ok = true
	return ret, nil
}

func (c *controlServer) GenerateAgentKey(ctx context.Context, req *spec.GenerateAgentKeyRequest) (*spec.GenerateAgentKeyResponse, error) {
	c.logger.Info("called handler")
	k, err := c.keyman.GenerateKey()
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "random number generation failed")
	}
	return &spec.GenerateAgentKeyResponse{AgentKey: k}, nil
}

func (c *controlServer) ListAgentKey(context.Context, *spec.ListAgentKeyRequest) (*spec.ListAgentKeyResponse, error) {
	d := c.keyman.GetKeysDigest()
	return &spec.ListAgentKeyResponse{Keys: d}, nil
}

func (c *controlServer) RevokeAgentKey(ctx context.Context, req *spec.RevokeAgentKeyRequest) (*spec.RevokeAgentKeyResponse, error) {
	if err := c.keyman.RemoveKey(req.GetId()); err != nil {
		return nil, status.Errorf(codes.Aborted, "revokation failed, maybe specified key is not exists")
	}
	return &spec.RevokeAgentKeyResponse{Ok: true}, nil
}

func loadAndSanitizeArgs() {
	flag.Parse()
	if _, _, err := net.ParseCIDR(*wireguardAddressRange); err != nil {
		log.Fatalf("unrecognized wireguardAddressRange, err=%s", err)
	}
	for _, s := range strings.Split(*assignAddressRange, ",") {
		if _, _, err := net.ParseCIDR(s); err != nil {
			log.Fatalf("unrecognized assignAddressRange, err=%s", err)
		}
	}
	if *adminPassword == "" {
		log.Fatalln("please set admin password with -adminPassword")
	}
}

type authorizer struct {
	password string
	server   *controlServer
	logger   *zap.Logger
}

func (a *authorizer) passwordAuth(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: no context metadata")
	}
	if len(md.Get("Bearer")) != 1 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Invalid Bearer token, check you send Bearer just one")
	}
	if md.Get("Bearer")[0] != a.password {
		a.logger.Info("auth failed", zap.String("got", md.Get("Bearer")[0]), zap.String("expected", a.password))
		return nil, grpc.Errorf(codes.Unauthenticated, "Wrong Bearer maybe password invalid")
	}
	return ctx, nil
}

func (a *authorizer) tokenAuth(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: no context metadata")
	}
	if len(md.Get("Bearer")) != 1 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Invalid Bearer token, check you send Bearer just one")
	}
	if err := a.server.keyman.IsValid(md.Get("Bearer")[0]); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "Wrong Bearer, maybe token invalid")
	}
	return ctx, nil
}

func (a *authorizer) HandleUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch {
	case strings.HasPrefix(info.FullMethod, "/ControlPlaneService/"):
		ctx, err := a.tokenAuth(ctx, info)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	case strings.HasPrefix(info.FullMethod, "/AdministratorService/"):
		a.logger.Info("pref admin")
		ctx, err := a.passwordAuth(ctx, info)
		if err != nil {
			return nil, err
		}
		a.logger.Info("auth ok")
		return handler(ctx, req)
	default:
		return nil, grpc.Errorf(codes.NotFound, "no such endpoint")
	}
}

func main() {
	ip.SetSeed(time.Now().UnixNano())
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	loadAndSanitizeArgs()
	server := controlServer{}
	server.stateFilePath = filepath.Join(*stateDir, "clients")
	if err := server.LoadState(); err != nil {
		logger.Panic("failed to load state", zap.Error(err))
	}
	server.logger = logger.Named("handler")
	a := authorizer{password: *adminPassword, server: &server, logger: logger.Named("authorizer")}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(a.HandleUnary),
	)
	spec.RegisterControlPlaneServiceServer(s, &server)
	spec.RegisterAdministratorServiceServer(s, &server)
	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		logger.Panic("failed to listen", zap.Error(err))
	}
	logger.Info("server started")
	if err := s.Serve(lis); err != nil {
		logger.Panic("failed to start serving", zap.Error(err))
	}
}
