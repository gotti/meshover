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
	"syscall"
	"time"

	"github.com/gotti/meshover/frr"
	"github.com/gotti/meshover/gre"
	"github.com/gotti/meshover/grpcclient"
	"github.com/gotti/meshover/linuxwireguard"
	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
)

//go:embed embed/daemons
var frrDaemonsConfig string

//go:embed embed/vtysh.conf
var frrVtyshConfig string

var (
	controlserver = flag.String("controlserver", "", "localhost:8080")
	hostName      = flag.String("hostname", "", "hostname")
	capabl        = flag.String("cap", "", "wireguard,linuxkernelwireguard")
	staticRoutes  = flag.String("static", "", "192.168.0.0/16")
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

func parseArgs() {
	if *hostName == "" {
		h, err := os.Hostname()
		if err != nil {
			log.Fatalf("failed to get hostname, err=%s", err)
		}
		hostName = &h
	}
}

func readConfig(confPath string) string {
	b, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalln("failed to read config")
	}
	return string(b)
}

func main() {
	flag.Parse()
	parseArgs()
	if *controlserver == "" {
		log.Fatalln("controlserver is not specified")
	}
	p := &spec.Peers{}
	stat := status.NewClient(*hostName, "meshover0", p)

	tunnel := linuxwireguard.NewTunnel(stat)
	defer func() {
		err := tunnel.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	addrs, err := getMachineAddresses()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	client, asn, err := grpcclient.NewClient(ctx, *hostName, *controlserver, tunnel.GetPublicKey(), &spec.AddressAndPort{Ipaddress: spec.NewAddress(addrs[0].IP), Port: 12912})
	defer client.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tunnel.SetAddressWithoutTouchingLinuxNetwork(client.OverlayIP)
	tunnel.SetListenPort(12912)
	stat.IPAddr = net.IPNet{IP: client.OverlayIP, Mask: net.IPv4Mask(255, 255, 255, 255)}
	stat.ASN = asn

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)
	watchctx, watchcancel := context.WithCancel(context.Background())
	defer watchcancel()

	c := make(chan []status.FrrPeer)

	greInstance := gre.NewGreInstance(stat.IPAddr.IP)
	defer greInstance.Clean()

	go func() {
		for {
			select {
			case <-ticker.C:
				res, err := client.ListPeers()
				if err != nil {
					log.Printf("failed to ListPeers err=%s", err)
				}
				pe := res.GetPeers()
				diff, err := stat.UpdatePeers(status.NewPeers(pe))
				if err != nil {
					log.Println(err)
				}
				if len(diff) > 0 {
					fmt.Println("updating...", diff)
					tunnel.UpdatePeers(diff)
					greInstance.UpdatePeers(diff)
					frrpeers := []status.FrrPeer{}
					for _, p := range pe.GetPeers() {
						fp := status.FrrPeer{Peer: p, TunName: greInstance.FindTunNameByOppositeIP(p.GetAddress().ToNetIP())}
						frrpeers = append(frrpeers, fp)
					}
					c <- frrpeers
				}
			case <-watchctx.Done():
				return
			}
		}
	}()
	conf := frr.NewFrrConfig(*hostName, client.OverlayIP.String(), readConfig("./conf/frr.conf"), frrDaemonsConfig, frrVtyshConfig)
	f, err := frr.NewInstance(ctx, asn, conf)
	if err != nil {
		log.Fatalln("failed to create frr instance", err)
	}
	defer func() {
		err := f.Clean()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	go func() {
		if f.Up(ctx); err != nil {
			log.Fatalln("failed to up frr instance", err)
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
