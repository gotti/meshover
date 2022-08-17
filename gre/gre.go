package gre

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
)

type GreStatus struct {
	counter int
	selfIP  net.IP
	tunnels []*GreTunnel
}

type GreTunnel struct {
	TunName    string
	link       netlink.Link
	peerName   string
	OppositeIP net.IP
}

var (
	errFileExists = fmt.Errorf("tunnel file exists")
)

func NewGreInstance(selfIP net.IP) *GreStatus {
	return &GreStatus{selfIP: selfIP}
}

func (g *GreStatus) addTunnelPeer(p *status.PeerDiffrence) error {
	tun := fmt.Sprintf("meshover0-tun%d", g.counter)
	attr := netlink.NewLinkAttrs()
	attr.Name = tun
	gretun := &netlink.Gretun{
		LinkAttrs: attr,
		Local:     g.selfIP,
		Remote:    p.Peer.GetAddress().ToNetIP(),
	}
	fmt.Printf("local=%s, remote=%s, device=%s\n", g.selfIP, p.Peer.GetAddress().ToNetIP(), gretun.LinkAttrs.Name)
	if err := netlink.LinkAdd(gretun); err != nil {
		if strings.Contains(err.Error(), "file exists") {
			if err := netlink.LinkDel(gretun); err != nil {
				log.Fatalln("failed to del", err)
			}
			return errFileExists
		}
		log.Fatalln("failed to create gre tunnel", err)
	}
	if err := netlink.LinkSetUp(gretun); err != nil {
		log.Fatalln("failed to up gre tunnel", err)
	}
	gt := &GreTunnel{TunName: tun, link: gretun, peerName: p.Peer.GetName(), OppositeIP: p.Peer.GetAddress().ToNetIP()}
	g.tunnels = append(g.tunnels, gt)
	fmt.Println("tunnels")
	for _, t := range g.tunnels {
		fmt.Printf("%+v", t)
	}
	g.counter++
	return nil
}

func Clean() {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalln("failed to delete gre peer", err)
	}
	for _, l := range links {
		if strings.HasPrefix(l.Attrs().Name, "meshover0-tun") {
			err := netlink.LinkDel(l)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func (g *GreStatus) delTunnelPeer(p *status.PeerDiffrence) {
	for i, r := range g.tunnels {
		if r.peerName == p.Peer.GetName() {
			if err := netlink.LinkDel(r.link); err != nil {
				log.Fatalln("failed to delete gre peer")
			}
			g.tunnels = append(g.tunnels[:i], g.tunnels[i+1:]...)
		}
	}
}

func (g *GreStatus) UpdatePeers(peersDiff []status.PeerDiffrence) {
	for _, p := range peersDiff {
		if p.Add {
			for {
				err := g.addTunnelPeer(&p)
				if err != nil {
					if err == errFileExists {
						fmt.Println("tunnel file already exists, try re creating after remove", err)
						continue
					} else {
						log.Fatalln("failed to add tunnel peer", err)
					}
				}
				break
			}
		} else {
			g.delTunnelPeer(&p)
		}
	}
}

func (g *GreStatus) FindTunNameByOppositeIP(opossiteIP net.IP) (tunName string) {
	for _, p := range g.tunnels {
		if p.OppositeIP.Equal(opossiteIP) {
			return p.TunName
		}
	}
	return ""
}
