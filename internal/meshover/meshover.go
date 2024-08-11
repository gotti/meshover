package meshover

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gotti/meshover/dummy"
	"github.com/gotti/meshover/frr"
	"github.com/gotti/meshover/gre"
	"github.com/gotti/meshover/grpcclient"
	"github.com/gotti/meshover/iproute"
	"github.com/gotti/meshover/kernel"
	"github.com/gotti/meshover/pkg/statuspusher"
	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"github.com/gotti/meshover/tunnel"
	"go.uber.org/zap"
)

// Settings contains meshover cli options and settings
type Settings struct {
	HostName        string
	ControlServer   string
	StatusServer    string
	AgentToken      string
	UnderlayIP      *net.IPNet
	RouteGathering  []*net.IPNet
	FrrVtyshConfig  string
	FrrDaemonConfig string
	FrrBGPConfig    string
	FrrBackend      frr.BackendType
	KernelSetting   *kernel.Settings
}

// Run runs meshover
func Run(ctx context.Context, logger *zap.Logger, settings Settings) error {
	if settings.KernelSetting == nil {
		kernel.NewDummyInstance()
	}
	if _, err := kernel.NewKernelInstance(settings.KernelSetting); err != nil {
		logger.Panic("failed to create new kernel instance", zap.Error(err))
	}
	stat := status.NewClient(settings.HostName, "meshover0", &spec.Peers{})
	//generate wireguard instance
	//generate keypair
	tun := tunnel.NewLinuxTunnel(logger, stat)
	defer func() {
		err := tun.Close()
		if err != nil {
			logger.Error("failed to close wireguard tunnel", zap.Error(err))
		}
	}()

	//send query for getting meshover ip
	client, asn, err := grpcclient.NewClient(ctx, logger, settings.HostName, settings.ControlServer, settings.AgentToken, tun.GetPublicKey(), &spec.AddressAndPort{Ipaddress: spec.NewAddress(settings.UnderlayIP.IP), Port: 12912}, settings.RouteGathering)
	if err != nil {
		return fmt.Errorf("failed to create new grpc connection, err=%w", err)
	}
	defer client.Close()

	//set got meshover ip before
	if err := tun.SetAddress(*client.OverlayIP); err != nil {
		return fmt.Errorf("failed to set tunnel address, err=%w", err)
	}
	tun.SetListenPort(12912)
	stat.IPAddr = *client.OverlayIP
	stat.ASN = asn

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)
	watchctx, watchcancel := context.WithCancel(ctx)
	defer watchcancel()

	c := make(chan []status.FrrPeerDiffrence)

	gre.Clean()
	greInstance := gre.NewGreInstance(logger, stat.IPAddr, *client.AdditionalIPs[0])
	defer gre.Clean()

	conf := frr.NewFrrConfig(settings.HostName, client.AdditionalIPs[0].IP.String(), settings.FrrBGPConfig, settings.FrrDaemonConfig, settings.FrrVtyshConfig)

	var f frr.Backend

	switch settings.FrrBackend {
	case frr.BackendNone:
		f = frr.NewDummyInstance()
	case frr.BackendDockerSDK:
		f, err = frr.NewDockerInstance(ctx, asn, conf)
		if err != nil {
			return fmt.Errorf("failed to create frr instance, err=%w", err)
		}
	case frr.BackendNerdCtl:
		f, err = frr.NewNerdCtlInstance(ctx, asn, conf)
		if err != nil {
			return fmt.Errorf("failed to create frr instance, err=%w", err)
		}
	default:
		return fmt.Errorf("no such backend")
	}
	defer func() {
		err := f.Kill()
		if err != nil {
			logger.Error("failed to cleanup frr container", zap.Error(err))
		}
	}()
	go func() {
		if err := f.Run(ctx); err != nil {
			logger.Fatal("failed to up frr instance", zap.Error(err))
		}
	}()
	iproute.ForceClean()
	iprouteInstance := iproute.NewInstance(logger)
	defer func() {
		if err := iprouteInstance.Clean(); err != nil {
			logger.Error("failed to clean iproute instance", zap.Error(err))
		}
	}()

	loopbackInterface, err := dummy.NewDevice("dummy-meshover0")
	if err != nil {
		return fmt.Errorf("failed to create new dummy interface, err=%w", err)
	}
	for i := range client.AdditionalIPs {
		loopbackInterface.AddAddress(client.AdditionalIPs[i])
	}
	defer func() {
		err := loopbackInterface.Clean()
		if err != nil {
			logger.Fatal("failed to clean loopbackInterface", zap.Error(err))
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
						oppsite, err := greInstance.FindTunNameByOppositeIP(d.NewPeer.GetWireguardAddress().ToNetIPNet().IP)
						fmt.Println("opossite", oppsite)
						if err != nil {
							log.Printf("failed to find tunname from oppositeIP \"%s\", err=%s\n", d.NewPeer.GetWireguardAddress().ToNetIPNet().IP, err)
						}
						fmt.Println("opposite", oppsite)
						frrdiff = append(frrdiff, status.FrrPeerDiffrence{PeerDiffrence: diff[i], TunName: oppsite})
					}
					sbrdiffs := []iproute.SBRDiff{}
					for _, d := range diff {
						oppsite, err := greInstance.FindTunNameByOppositeIP(d.NewPeer.GetWireguardAddress().ToNetIPNet().IP)
						if err != nil {
							log.Printf("failed to find tunname from oppositeIP \"%s\", err=%s\n", d.NewPeer.GetWireguardAddress(), err)
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
					if err := tun.UpdatePeers(diff); err != nil {
						logger.Error("failed to update tunnel peer", zap.Error(err))
					}
					c <- frrdiff
				}
			case <-watchctx.Done():
				return
			}
		}
	}()
	go func() {
		if settings.StatusServer == "" {
			return
		}
		statusClient, err := statuspusher.NewClient(ctx, logger, settings.StatusServer)
		if err != nil {
			log.Fatalf("failed to create status pusher, err=%s", err)
		}
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				s := receiveFromChan(logger, statusBGP)
				req := &spec.RegisterStatusRequest{
					Status: &spec.StatusManagerPeerStatus{
						Hostname: settings.HostName,
						NodeStatus: &spec.MinimumNodeStatus{
							LocalAS: asn,
							Addresses: []*spec.AddressCIDR{
								spec.NewAddressCIDR(*client.AdditionalIPs[0]),
							},
							Endpoint: spec.NewAddress(settings.UnderlayIP.IP),
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
	return nil
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
