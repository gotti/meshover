package tunnel

import (
	//"encoding/binary"
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
	"go.uber.org/zap"
)

// LinuxTunnel is a wireguard tunnel device
type LinuxTunnel struct {
	logger  *zap.Logger
	selfIP  net.IP
	link    netlink.Link
	keyPair *spec.Curve25519KeyPair
}

// GetPublicKey get public key of tunnel
func (t *LinuxTunnel) GetPublicKey() *spec.Curve25519Key {
	return t.keyPair.GetPublickey()
}

// NewLinuxTunnel creates wiregard device
func NewLinuxTunnel(logger *zap.Logger, peers *status.ClientStatus) LinuxTunnel {
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
	tun := LinuxTunnel{keyPair: keypair, link: la, selfIP: peers.IPAddr.IP}
	tun.setPrivateKey(keypair.GetPrivatekey())
	return tun
}

// Up set link status to up
func (t *LinuxTunnel) Up() error {
	if err := netlink.LinkSetUp(t.link); err != nil {
		return fmt.Errorf("failed to up device, err=%w", err)
	}
	return nil
}

func (t *LinuxTunnel) setPrivateKey(k *spec.Curve25519Key) error {
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

// SetListenPort sets wireguard underlay listen port
func (t *LinuxTunnel) SetListenPort(port int) error {
	err := exec.Command("wg", "set", t.link.Attrs().Name, "listen-port", fmt.Sprint(port)).Run()
	if err != nil {
		return fmt.Errorf("failed to set listenport, err=%w", err)
	}
	return nil
}

// Close deletes link
func (t *LinuxTunnel) Close() error {
	err := netlink.LinkDel(t.link)
	if err != nil {
		return fmt.Errorf("failed to delete link, err=%w", err)
	}
	return nil
}

func (t *LinuxTunnel) addPeer(p *spec.Peer) error {
	u := p.GetUnderlayLinuxKernelWireguard()
	o, err := exec.Command("wg", "set", t.link.Attrs().Name, "peer", u.GetPublicKey().EncodeBase64(), "allowed-ips", p.GetWireguardAddress().Format(), "endpoint", u.GetEndpoint().Format(), "persistent-keepalive", "10").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set peer, out=%s, err=%w", string(o), err)
	}
	return nil
}

func (t *LinuxTunnel) delPeer(p *spec.Peer) error {
	u := p.GetUnderlayLinuxKernelWireguard()
	o, err := exec.Command("wg", "set", t.link.Attrs().Name, "peer", u.GetPublicKey().EncodeBase64(), "remove").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to del peer, out=%s err=%w", string(o), err)
	}
	return nil
}

func (t *LinuxTunnel) addRoute(p *spec.Peer) error {
	_, n, err := net.ParseCIDR(p.GetWireguardAddress().ToNetIPNet().String())
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

func (t *LinuxTunnel) delRoute(p *spec.Peer) error {
	_, n, err := net.ParseCIDR(p.GetWireguardAddress().ToNetIPNet().String())
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

// UpdatePeers updates peer
func (t *LinuxTunnel) UpdatePeers(peersDiff []status.PeerDiffrence) error {
	return updatePeers(t.logger, t, peersDiff)
}

/*
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
*/

// SetAddress sets wireguard address of oneself
func (t *LinuxTunnel) SetAddress(address net.IPNet) error {
	addr, err := netlink.ParseAddr(address.String())
	if err != nil {
		return fmt.Errorf("parseaddr, err=%w", err)
	}
	netlink.AddrAdd(t.link, addr)
	t.selfIP = address.IP
	netlink.LinkSetUp(t.link)
	return nil
}
