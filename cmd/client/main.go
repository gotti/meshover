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
	hostName          = flag.String("hostname", "", "hostname")
	capabl            = flag.String("cap", "", "wireguard,linuxkernelwireguard")
	staticRoutes      = flag.String("static", "", "192.168.0.0/16")
	rawRouteGathering = flag.String("gathering", "", "1.1.1.0/27,1.1.2.0/29")
)

var (
	routeGathering []*net.IPNet = nil
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
	if err := tunnel.SetAddress(client.OverlayIP); err != nil {
		panicError(logger, "failed to set address", err)
	}
	tunnel.SetListenPort(12912)
	stat.IPAddr = net.IPNet{IP: client.OverlayIP.IP, Mask: net.IPv4Mask(255, 255, 255, 255)}
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

	conf := frr.NewFrrConfig(*hostName, client.OverlayIP.IP.String(), readConfig("./conf/frr.conf"), frrDaemonsConfig, frrVtyshConfig)
	f, err := frr.NewInstance(ctx, asn, conf)
	if err != nil {
		panicError(logger, "failed to create frr instance", err)
	}
	defer func() {
		err := f.Clean()
		if err != nil {
			panicError(logger, "failed to cleanup frr container", err)
		}
	}()
	go func() {
		if f.Up(ctx); err != nil {
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
		for {
			select {
			case p := <-c:
				if err := f.UpdatePeers(ctx, p); err != nil {
					log.Fatalln("failed to update AS peer", err)
				}
			}
		}
	}()

	select {
	case <-term:
	}
}
