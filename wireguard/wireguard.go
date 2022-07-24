package wireguard

import (
	"fmt"
	"log"
	"net"

	"github.com/gotti/meshover/spec"
	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
)

type WireguardTunnel struct {
  hostName string
  ipAddr    net.IP
	tunName    string
	privatekey []byte
	publickey  []byte
	peers      *spec.Peers
	Dev        *device.Device
  link netlink.Link
}

func NewTunnel(tunName string, privatekey, publickey []byte, hostName string) WireguardTunnel {
	t, err := tun.CreateTUN(tunName, 1420)
	if err != nil {
		log.Panic(err)
	}
	la, err := netlink.LinkByName(tunName)
	if err != nil {
		log.Fatalln(err)
	}
	dev := device.NewDevice(t, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	fmt.Println(dev.IpcGet())
  return WireguardTunnel{hostName: hostName, tunName: tunName, privatekey: privatekey, publickey: publickey, Dev: dev, link: la}
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

func (t *WireguardTunnel) SetPeers(peers *spec.Peers) {
	t.peers = peers
	var conf string
	conf += fmt.Sprintf("private_key=%x\n", t.privatekey)
	conf += fmt.Sprintf("listen_port=12912\n")
  if t.peers == nil {
    return
  }
	for _, p := range t.peers.Peers {
		u := p.GetUnderlayWireguard()
		if u == nil {
			log.Printf("ignoring %s\n", u.Endpoint)
			continue
		}
    if p.GetName() == t.hostName {
      log.Printf("%s is myself, skipping\n", p.GetName())
      continue
    }
		conf += fmt.Sprintf("public_key=%x\n", u.GetPublicKey())
		conf += fmt.Sprintf("endpoint=%s\n", u.GetEndpoint().GetEndpointAddress())
    conf += fmt.Sprintf("allowed_ip=%s/32\n", p.GetAddress().GetEndpointAddress())
		conf += fmt.Sprintf("persistent_keepalive_interval=10")
    route := new(netlink.Route)
    _, n, err := net.ParseCIDR(fmt.Sprintf("%s/32", p.GetAddress().GetEndpointAddress()))
    if err != nil {
      log.Printf("failed to parse cidr, err=%s\n", err)
    }
    route.Dst = n
    route.LinkIndex = t.link.Attrs().Index
    if err := netlink.RouteAdd(route); err != nil {
      log.Printf("failed to add route to %s, err=%s", p.GetAddress().GetEndpointAddress(), err)
    }
	}
	t.setIpc(conf)
}

func (t *WireguardTunnel) SetAddress(address net.IP) {
  t.ipAddr = address
	addr, err := netlink.ParseAddr(fmt.Sprintf("%s/32", address.String()))
	if err != nil {
		log.Fatalln("parseaddr", err)
	}
	netlink.AddrAdd(t.link, addr)
	netlink.LinkSetUp(t.link)
}
