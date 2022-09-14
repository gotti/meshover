package tunnel

import (
	"fmt"

	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
	"go.uber.org/zap"
)

// Tunnel is a interface that interconnect each node
type Tunnel interface {
	UpdatePeers(peerdiff []status.PeerDiffrence) error
	addPeer(peer *spec.Peer) error
	addRoute(peer *spec.Peer) error
	delPeer(peer *spec.Peer) error
	delRoute(peer *spec.Peer) error
}

// updatePeers updates tunnel peer
func updatePeers(logger *zap.Logger, t Tunnel, peersDiff []status.PeerDiffrence) error {
	for _, p := range peersDiff {
		if ca := p.NewPeer.GetUnderlayLinuxKernelWireguard(); ca == nil {
			logger.Info("unknown underlay", zap.String("underlay", p.NewPeer.String()))
			continue
		}
		switch p.Diff {
		case status.DiffTypeAdd:
			{
				if err := t.addPeer(p.NewPeer); err != nil {
					return fmt.Errorf("ADD: failed to add peer, err=%s", err)
				}
				if err := t.addRoute(p.NewPeer); err != nil {
					return fmt.Errorf("ADD: failed to add route, err=%s", err)
				}
			}
		case status.DiffTypeDelete:
			{
				if err := t.delPeer(p.OldPeer); err != nil {
					return fmt.Errorf("DEL: failed to add peer, err=%s", err)
				}
				if err := t.delRoute(p.OldPeer); err != nil {
					return fmt.Errorf("DEL: failed to add route, err=%s", err)
				}
			}
		case status.DiffTypeChange:
			{
				if err := t.delPeer(p.OldPeer); err != nil {
					return fmt.Errorf("CHG: failed to add peer, err=%s", err)
				}
				if err := t.delRoute(p.OldPeer); err != nil {
					return fmt.Errorf("CHG: failed to add route, err=%s", err)
				}
				if err := t.addPeer(p.NewPeer); err != nil {
					return fmt.Errorf("CHG: failed to add peer, err=%s", err)
				}
				if err := t.addRoute(p.NewPeer); err != nil {
					return fmt.Errorf("CHG: failed to add route, err=%s", err)
				}
			}
		}
	}
	return nil
}
