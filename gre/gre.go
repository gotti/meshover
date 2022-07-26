package gre

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

type GreStatus struct {
	log        *zap.Logger
	counter    int
	underlayIP net.IPNet
	overlayIP  net.IPNet
	tunnels    []*GreTunnel
}

type GreTunnel struct {
	TunName    string
	link       netlink.Link
	peerName   string
	OppositeIP net.IPNet
}

var (
	errFileExists = fmt.Errorf("tunnel file exists")
)

func NewGreInstance(log *zap.Logger, underlayIP, overlayIP net.IPNet) *GreStatus {
	return &GreStatus{log: log, underlayIP: underlayIP, overlayIP: overlayIP}
}

func addressCIDRArrayToIPNetArray(a []*spec.AddressCIDR) []net.IPNet {
	var ret []net.IPNet
	for _, d := range a {
		ret = append(ret, *d.ToNetIPNet())
	}
	return ret
}

func (g *GreStatus) addTunnelPeer(p *status.PeerDiffrence) error {
	tun := fmt.Sprintf("meshover0-tun%d", g.counter)
	attr := netlink.NewLinkAttrs()
	attr.Name = tun
	gretun := &netlink.Gretun{
		LinkAttrs: attr,
		Local:     g.underlayIP.IP,
		Remote:    p.NewPeer.GetWireguardAddress().ToNetIPNet().IP,
	}
	g.log.Debug("connection info", zap.String("self", g.underlayIP.String()), zap.String("target", p.NewPeer.GetWireguardAddress().Format()), zap.String("tun", gretun.LinkAttrs.Name))
	if err := netlink.LinkAdd(gretun); err != nil {
		if strings.Contains(err.Error(), "file exists") {
			if err := netlink.LinkDel(gretun); err != nil {
				log.Fatalln("failed to del", err)
			}
			return errFileExists
		}
		log.Fatalln("failed to create gre tunnel", err)
	}
	if err := netlink.AddrAdd(gretun, &netlink.Addr{IPNet: &g.overlayIP, Label: ""}); err != nil {
		log.Fatalln("failed to set address", err)
	}
	if err := netlink.LinkSetUp(gretun); err != nil {
		log.Fatalln("failed to up gre tunnel", err)
	}
	gt := &GreTunnel{TunName: tun, link: gretun, peerName: p.NewPeer.GetName(), OppositeIP: *p.NewPeer.WireguardAddress.ToNetIPNet()}
	g.tunnels = append(g.tunnels, gt)
	fmt.Println("tunnels")
	for _, t := range g.tunnels {
		fmt.Printf("%+v", t)
	}
	g.counter++
	return nil
}

func (g *GreStatus) searchByHostName(hostName string) *GreTunnel {
	for _, d := range g.tunnels {
		if d.peerName == hostName {
			return d
		}
	}
	return nil
}

func (g *GreStatus) updateTunnelPeer(p *status.PeerDiffrence) error {
	searched := g.searchByHostName(p.NewPeer.GetName())
	if searched == nil {
		err := fmt.Errorf("update failed, not found such peer, peerName=%s", p.NewPeer.GetName())
		g.log.Error("update faled", zap.Error(err))
		return err
	}
	attr := netlink.NewLinkAttrs()
	attr.Name = searched.TunName
	gretun := &netlink.Gretun{
		LinkAttrs: attr,
		Local:     g.underlayIP.IP,
		Remote:    p.NewPeer.GetWireguardAddress().ToNetIPNet().IP,
	}
	if err := netlink.LinkModify(gretun); err != nil {
		return fmt.Errorf("failed to modify link, err=%w", err)
	}
	searched.link = gretun
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
		if r.peerName == p.OldPeer.GetName() {
			if err := netlink.LinkDel(r.link); err != nil {
				log.Fatalln("failed to delete gre peer")
			}
			g.tunnels = append(g.tunnels[:i], g.tunnels[i+1:]...)
		}
	}
}

func (g *GreStatus) UpdatePeers(peersDiff []status.PeerDiffrence) {
	for _, p := range peersDiff {
		switch p.Diff {
		case status.DiffTypeAdd:
			{
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
			}
		case status.DiffTypeDelete:
			{
				g.delTunnelPeer(&p)
			}
		case status.DiffTypeChange:
			{
				g.delTunnelPeer(&p)
				if err := g.addTunnelPeer(&p); err != nil {
					log.Println("@@@@@@@@@@@@ failed to add tunnel peer", err)
				}
			}
		}
	}
}

func (g *GreStatus) FindTunNameByOppositeIP(opossiteIP net.IP) (tunName string, err error) {
	for _, p := range g.tunnels {
		if p.OppositeIP.IP.Equal(opossiteIP) {
			return p.TunName, nil
		}
	}
	return "", fmt.Errorf("not found such opposite IP")
}
