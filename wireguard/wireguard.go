package wireguard

import (
	"fmt"
	"log"
	"net"

	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
)

type WireguardTunnel struct {
	Dev        *device.Device
	link       netlink.Link
	privatekey string
	publickey  string
}

func NewTunnel(peers *status.ClientStatus, privatekey, publickey string) WireguardTunnel {
	t, err := tun.CreateTUN(peers.TunName, 1420)
	if err != nil {
		log.Panic(err)
	}
	la, err := netlink.LinkByName(peers.TunName)
	if err != nil {
		log.Fatalln(err)
	}
	dev := device.NewDevice(t, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	var conf string
	conf += fmt.Sprintf("private_key=%x\n", privatekey)
	conf += fmt.Sprintf("listen_port=12912\n")
	conf += ""
	dev.IpcSet(conf)
	return WireguardTunnel{privatekey: privatekey, publickey: publickey, Dev: dev, link: la}
}

func (t *WireguardTunnel) Up() error {
	if err := t.Dev.Up(); err != nil {
		return fmt.Errorf("failed to up device, err=%w", err)
	}
	return nil
}

func (t *WireguardTunnel) Close() {
	t.Dev.Close()
}

func (t *WireguardTunnel) setIpc(ipc string) {
	t.Dev.IpcSet(ipc)
}

func (t *WireguardTunnel) UpdatePeers(peersDiff []status.PeerDiffrence) {
	var conf string
	for _, p := range peersDiff {
		ca := p.NewPeer.GetUnderlayLinuxKernelWireguard()
		if ca == nil {
			continue
		}
		conf += fmt.Sprintf("public_key=%x\n", ca.GetPublicKey())
		conf += fmt.Sprintf("endpoint=%s\n", ca.GetEndpoint().GetEndpointAddress())
		conf += fmt.Sprintf("allowed_ip=%s/32\n", p.NewPeer.GetAddress().GetEndpointAddress())
		conf += fmt.Sprintf("persistent_keepalive_interval=10\n")
		if !p.Add {
			conf += fmt.Sprintf("remove=true\n")
		}
		route := new(netlink.Route)
		_, n, err := net.ParseCIDR(fmt.Sprintf("%s/32", p.NewPeer.GetAddress().GetEndpointAddress()))
		if err != nil {
			log.Printf("failed to parse cidr, err=%s\n", err)
		}
		route.Dst = n
		route.LinkIndex = t.link.Attrs().Index
		if err := netlink.RouteAdd(route); err != nil {
			log.Printf("failed to add route to %s, err=%s", p.NewPeer.GetAddress().GetEndpointAddress(), err)
		}
	}
	fmt.Println("applied", conf)
	t.setIpc(conf)
}

func (t *WireguardTunnel) SetAddress(address net.IP) {
	addr, err := netlink.ParseAddr(fmt.Sprintf("%s/32", address.String()))
	if err != nil {
		log.Fatalln("parseaddr", err)
	}
	netlink.AddrAdd(t.link, addr)
	netlink.LinkSetUp(t.link)
}
