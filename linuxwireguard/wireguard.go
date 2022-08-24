package linuxwireguard

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"github.com/vishvananda/netlink"
)

type WireguardTunnel struct {
	selfIP  net.IP
	link    netlink.Link
	keyPair *spec.Curve25519KeyPair
}

func (t *WireguardTunnel) GetPublicKey() *spec.Curve25519Key {
	return t.keyPair.GetPublickey()
}

func NewTunnel(peers *status.ClientStatus) WireguardTunnel {
	laa := netlink.NewLinkAttrs()
	laa.Name = peers.TunName
	wg := &netlink.Wireguard{LinkAttrs: laa}
	err := netlink.LinkAdd(wg)
	if err != nil {
		if strings.Contains(err.Error(), "file exists") {
		} else {
			log.Panic(err)
		}
	}
	la, err := netlink.LinkByName(peers.TunName)
	if err != nil {
		log.Fatalln(err)
	}
	keypair, err := spec.GenerateKeyPair()
	if err != nil {
		log.Fatalln(err)
	}
	tun := WireguardTunnel{keyPair: keypair, link: la, selfIP: peers.IPAddr.IP}
	tun.setPrivateKey(keypair.GetPrivatekey())
	return tun
}

func (t *WireguardTunnel) Up() error {
	if err := netlink.LinkSetUp(t.link); err != nil {
		return fmt.Errorf("failed to up device, err=%w", err)
	}
	return nil
}

func (t *WireguardTunnel) setPrivateKey(k *spec.Curve25519Key) error {
	tmpName := filepath.Join(os.TempDir(), "priv")
	if err := os.WriteFile(tmpName, []byte(k.EncodeBase64()), 0700); err != nil {
		return fmt.Errorf("failed to write private key, err=%w", err)
	}
	defer os.Remove(tmpName)
	if err := exec.Command("wg", "set", t.link.Attrs().Name, "private-key", tmpName).Run(); err != nil {
		return fmt.Errorf("failed to set private key, err=%w", err)
	}
	return nil
}

func (t *WireguardTunnel) SetListenPort(port int) error {
	err := exec.Command("wg", "set", t.link.Attrs().Name, "listen-port", fmt.Sprint(port)).Run()
	if err != nil {
		return fmt.Errorf("failed to set listenport, err=%w", err)
	}
	return nil
}

func (t *WireguardTunnel) Close() error {
	err := netlink.LinkDel(t.link)
	if err != nil {
		return fmt.Errorf("failed to delete link, err=%w", err)
	}
	return nil
}

func (t *WireguardTunnel) setPeer(p *spec.Peer) error {
	u := p.GetUnderlayLinuxKernelWireguard()
	o, err := exec.Command("wg", "set", t.link.Attrs().Name, "peer", u.GetPublicKey().EncodeBase64(), "allowed-ips", p.GetAddress()[0].String(), "endpoint", fmt.Sprintf("%s", u.GetEndpoint().Format()), "persistent-keepalive", "10").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set peer, out=%s, err=%w", string(o), err)
	}
	return nil
}

func (t *WireguardTunnel) delPeer(p *spec.Peer) error {
	u := p.GetUnderlayLinuxKernelWireguard()
	o, err := exec.Command("wg", "set", t.link.Attrs().Name, "peer", u.GetPublicKey().EncodeBase64(), "remove").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to del peer, out=%s err=%w", string(o), err)
	}
	return nil
}

func (t *WireguardTunnel) addRoute(p *spec.Peer) error {
	_, n, err := net.ParseCIDR(p.GetAddress()[0].ToNetIPNet().String())
	if err != nil {
		log.Printf("failed to parse cidr, err=%s\n", err)
	}
	route := new(netlink.Route)
	route.Dst = n
	route.LinkIndex = t.link.Attrs().Index
	route.Src = t.selfIP
	fmt.Println("self", t.selfIP.String())
	route.Scope = netlink.SCOPE_LINK
	if err := netlink.RouteAdd(route); err != nil {
		return fmt.Errorf("failed to add route to %s, err=%w", p.GetAddress(), err)
	}
	return nil
}

func (t *WireguardTunnel) delRoute(p *spec.Peer) error {
	_, n, err := net.ParseCIDR(p.GetAddress()[0].ToNetIPNet().String())
	if err != nil {
		log.Printf("failed to parse cidr, err=%s\n", err)
	}
	route := new(netlink.Route)
	route.Dst = n
	route.LinkIndex = t.link.Attrs().Index
	if err := netlink.RouteDel(route); err != nil {
		return fmt.Errorf("failed to delete route to %s, err=%w", p.GetAddress(), err)
	}
	return nil
}

func (t *WireguardTunnel) UpdatePeers(peersDiff []status.PeerDiffrence) {
	for _, p := range peersDiff {
		if ca := p.Peer.GetUnderlayLinuxKernelWireguard(); ca == nil {
			fmt.Println("unknown underlay, skipping...", p)
			continue
		}
		if p.Add {
			fmt.Println("adding")
			if err := t.setPeer(p.Peer); err != nil {
				log.Println(err)
			}
			if err := t.addRoute(p.Peer); err != nil {
				log.Println(err)
			}
		} else {
			if err := t.delPeer(p.Peer); err != nil {
				log.Println(err)
			}
			if err := t.delRoute(p.Peer); err != nil {
				log.Println(err)
			}
		}
	}
}

func generateLinkLocalV6(address net.IP) net.IP {
	linklocalv6 := make(net.IP, 16)
	if v4 := address.To4(); v4 != nil {
		binary.BigEndian.PutUint32(linklocalv6[:4], 0xfe800000)
		binary.BigEndian.PutUint32(linklocalv6[4:8], 0xdeadbeef)
		binary.BigEndian.PutUint32(linklocalv6[8:12], 0xdeadbeef)
		for i := 0; i < 4; i++ {
			linklocalv6[i+12] = v4[i]
		}
		return linklocalv6
	}
	if v6 := address.To16(); v6 != nil {
		binary.BigEndian.PutUint32(linklocalv6[:4], 0xfe800000)
		binary.BigEndian.PutUint32(linklocalv6[4:8], 0xdeadbeef)
		copy(linklocalv6[9:], v6)
	}
	return linklocalv6
}

func (t *WireguardTunnel) SetAddress(address net.IP) {
	addr, err := netlink.ParseAddr(fmt.Sprintf("%s/32", address.String()))
	if err != nil {
		log.Fatalln("parseaddr", err)
	}
	netlink.AddrAdd(t.link, addr)
	t.selfIP = address
	netlink.LinkSetUp(t.link)
}

func (t *WireguardTunnel) SetAddressWithoutTouchingLinuxNetwork(address net.IP) {
	t.selfIP = address
}
