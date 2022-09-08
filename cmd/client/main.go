package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gotti/meshover/frr"
	"github.com/gotti/meshover/gre"
	"github.com/gotti/meshover/grpcclient"
	"github.com/gotti/meshover/iproute"
	"github.com/gotti/meshover/linuxwireguard"
	"github.com/gotti/meshover/pkg/statuspusher"
	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

//go:embed embed/daemons
var frrDaemonsConfig string

//go:embed embed/vtysh.conf
var frrVtyshConfig string

var (
	controlserver     = flag.String("controlserver", "", "localhost:8080")
	statusserver      = flag.String("statusserver", "", "localhost:8080")
	rawfrrBackend     = flag.String("frr", "", "select one of following: none, dockersdk, nerdctl")
	hostName          = flag.String("hostname", "", "hostname")
	capabl            = flag.String("cap", "", "wireguard,linuxkernelwireguard")
	staticRoutes      = flag.String("static", "", "192.168.0.0/16")
	rawRouteGathering = flag.String("gathering", "", "1.1.1.0/27,1.1.2.0/29")
)

type frrBackendType string

var (
	routeGathering      []*net.IPNet = nil
	frrBackendNone                   = frrBackendType("none")
	frrBackendDockerSDK              = frrBackendType("dockersdk")
	frrBackendNerdCtl                = frrBackendType("nerdctl")
	frrBackend                       = frrBackendNone
)

func getMachineAddresses() (ret []*net.IPNet, err error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for _, l := range links {
		if l.Type() != "device" && l.Type() != "vlan" {
			continue
		}
		addr, err := netlink.AddrList(l, syscall.AF_INET6)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch ip address, err=%w", err)
		}
		for _, a := range addr {
			if a.IP.IsGlobalUnicast() {
				ret = append(ret, a.IPNet)
			}
		}
	}
	return ret, nil
}

func parseArgs() error {
	if *hostName == "" {
		h, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("failed to get hostname, err=%w", err)
		}
		hostName = &h
	}
	if *rawRouteGathering != "" {
		for _, r := range strings.Split(*rawRouteGathering, ",") {
			_, n, err := net.ParseCIDR(r)
			if err != nil {
				return fmt.Errorf("failed to parse route gathering option, err=%w", err)
			}
			routeGathering = append(routeGathering, n)
		}
	}
	if *controlserver == "" {
		return fmt.Errorf("controlserver is not specified")
	}
	switch *rawfrrBackend {
	case "none":
		frrBackend = frrBackendNone
	case "dockersdk":
		frrBackend = frrBackendDockerSDK
	case "nerdctl":
		frrBackend = frrBackendNerdCtl
	default:
		return fmt.Errorf("please select valid frr backend with -frr")
	}
	return nil
}

func readConfig(confPath string) string {
	b, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalln("failed to read config")
	}
	return string(b)
}

func panicError(log *zap.Logger, message string, err error) {
	log.Panic(message, zap.Error(err))
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("failed to initialize zap")
	}
	flag.Parse()
	if err := parseArgs(); err != nil {
		panicError(logger, "failed to parse args", err)
	}
	stat := status.NewClient(*hostName, "meshover0", &spec.Peers{})

	//get ipv6 unicast address from device or vlan
	addrs, err := getMachineAddresses()
	if err != nil {
		panicError(logger, "failed to get ipv6 unicast address", err)
	}

	//generate wireguard instance
	//generate keypair
	tunnel := linuxwireguard.NewTunnel(stat)
	defer func() {
		err := tunnel.Close()
		if err != nil {
			panicError(logger, "failed to close wireguard tunnel", err)
		}
	}()

	ctx := context.Background()
	//send query for getting meshover ip
	client, asn, err := grpcclient.NewClient(ctx, logger, *hostName, *controlserver, tunnel.GetPublicKey(), &spec.AddressAndPort{Ipaddress: spec.NewAddress(addrs[0].IP), Port: 12912}, routeGathering)
	defer client.Close()
	if err != nil {
		panicError(logger, "failed to create new grpc connection", err)
	}

	//set got meshover ip before
	if err := tunnel.SetAddress(*client.OverlayIP[0]); err != nil {
		panicError(logger, "failed to set address", err)
	}
	tunnel.SetListenPort(12912)
	stat.IPAddr = net.IPNet{IP: client.OverlayIP[0].IP, Mask: net.IPv4Mask(255, 255, 255, 255)}
	stat.ASN = asn

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)
	watchctx, watchcancel := context.WithCancel(context.Background())
	defer watchcancel()

	c := make(chan []status.FrrPeerDiffrence)

	gre.Clean()
	greInstance := gre.NewGreInstance(logger, stat.IPAddr.IP)
	defer gre.Clean()

	conf := frr.NewFrrConfig(*hostName, client.OverlayIP[0].IP.String(), readConfig("./conf/frr.conf"), frrDaemonsConfig, frrVtyshConfig)

	var f frr.Backend

	switch frrBackend {
	case frrBackendNone:
		f = frr.NewDummyInstance()
	case frrBackendDockerSDK:
		f, err = frr.NewDockerInstance(ctx, asn, conf)
		if err != nil {
			panicError(logger, "failed to create frr instance", err)
		}
	case frrBackendNerdCtl:
		f, err = frr.NewNerdCtlInstance(ctx, asn, conf)
		if err != nil {
			panicError(logger, "failed to create frr instance", err)
		}
	}
	defer func() {
		err := f.Kill()
		if err != nil {
			panicError(logger, "failed to cleanup frr container", err)
		}
	}()
	go func() {
		if err := f.Run(ctx); err != nil {
			panicError(logger, "failed to up frr instance", err)
		}
	}()
	iproute.ForceClean()
	iprouteInstance := iproute.NewInstance(logger)
	defer func() {
		if err := iprouteInstance.Clean(); err != nil {
			panicError(logger, "failed to clean iproute instance", err)
		}
	}()

	statusBGP := make(chan []*spec.PeerBGPStatus, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				res, err := client.ListPeers()
				if err != nil {
					log.Printf("failed to ListPeers err=%s", err)
				}
				pe := res.GetPeers()
				ppp := status.NewPeers(pe)
				diff, err := stat.UpdatePeers(ppp)
				if err != nil {
					log.Println(err)
				}
				if len(diff) > 0 {
					frrdiff := []status.FrrPeerDiffrence{}
					greInstance.UpdatePeers(diff)
					for i, d := range diff {
						oppsite, err := greInstance.FindTunNameByOppositeIP(d.NewPeer.GetAddress()[0].ToNetIPNet().IP)
						fmt.Println("opossite", oppsite)
						if err != nil {
							log.Printf("failed to find tunname from oppositeIP \"%s\", err=%s\n", d.NewPeer.GetAddress()[0].ToNetIPNet(), err)
						}
						fmt.Println("opposite", oppsite)
						frrdiff = append(frrdiff, status.FrrPeerDiffrence{PeerDiffrence: diff[i], TunName: oppsite})
					}
					sbrdiffs := []iproute.SBRDiff{}
					for _, d := range diff {
						oppsite, err := greInstance.FindTunNameByOppositeIP(d.NewPeer.GetAddress()[0].ToNetIPNet().IP)
						if err != nil {
							log.Printf("failed to find tunname from oppositeIP \"%s\", err=%s\n", d.NewPeer.GetAddress()[0].ToNetIPNet(), err)
						}
						sbr := d.NewPeer.GetSbrOption()
						if sbr == nil {
							continue
						}
						sbrdiff := iproute.SBRDiff{
							HostName:  d.NewPeer.GetName(),
							Diff:      d.Diff,
							TunName:   oppsite,
							Gathering: []net.IPNet{},
						}
						for _, r := range sbr.GetSourceIPRange() {
							sbrdiff.Gathering = append(sbrdiff.Gathering, *r.ToNetIPNet())
						}
						sbrdiffs = append(sbrdiffs, sbrdiff)
					}
					iprouteInstance.UpdateSBRPeer(sbrdiffs)
					fmt.Println("updating...", frrdiff)
					tunnel.UpdatePeers(diff)
					c <- frrdiff
				}
			case <-watchctx.Done():
				return
			}
		}
	}()
	go func() {
		statusClient, err := statuspusher.NewClient(ctx, logger, *statusserver)
		if err != nil {
			panicError(logger, "failed to create statuspusher", err)
		}
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				s := receiveFromChan(logger, statusBGP)
				req := &spec.RegisterStatusRequest{
					Status: &spec.StatusManagerPeerStatus{
						Hostname: *hostName,
						NodeStatus: &spec.MinimumNodeStatus{
							LocalAS: asn,
							Addresses: []*spec.AddressCIDR{
								spec.NewAddressCIDR(*client.OverlayIP[0]),
							},
							Endpoint: spec.NewAddress(addrs[0].IP),
						},
						PeerStatus: s,
					},
				}
				if err := statusClient.RegisterStatus(ctx, req); err != nil {
					logger.Error("failed to register status", zap.Error(err))
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case p := <-c:
				if err := f.UpdatePeers(ctx, p); err != nil {
					log.Fatalln("failed to update AS peer", err)
				}
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				stat, err := frr.GetBGPStatus(ctx, f)
				if err != nil {
					logger.Error("failed to get bgp status", zap.Error(err))
				}
				select {
				case statusBGP <- stat:
					logger.Debug("bgp status updated")
				default:
				}
			}
		}
	}()

	select {
	case <-term:
	}
}

func receiveFromChan(logger *zap.Logger, bgpChan <-chan []*spec.PeerBGPStatus) []*spec.NodePeersStatus {
	ret := []*spec.NodePeersStatus{}
	remotesMap := map[string]*spec.NodePeersStatus{}
	select {
	case bg := <-bgpChan:
		{
			for _, b := range bg {
				_, ok := remotesMap[b.GetRemoteHostname()]
				if !ok {
					remotesMap[b.GetRemoteHostname()] = &spec.NodePeersStatus{RemoteHostname: b.GetRemoteHostname(), BgpStatus: b}
				}
				remotesMap[b.GetRemoteHostname()].BgpStatus = b
			}
		}
	default:
		logger.Debug("bgp status not prepared")
		return nil
	}
	for _, v := range remotesMap {
		ret = append(ret, v)
	}
	return ret
}
