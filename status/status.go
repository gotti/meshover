package status

import (
	"fmt"
	"net"
	"sync"

	"github.com/gotti/meshover/spec"
	"google.golang.org/protobuf/proto"
)

type FrrPeerDiffrence struct {
	TunName string
	PeerDiffrence
}

// DiffType is a peer changes, enum of Add,Delete,Change
type DiffType string

const (
	//DiffTypeAdd shows this node is added and was not exist before
	DiffTypeAdd = DiffType("Add")
	//DiffTypeDelete shows this node is deleted and existed before
	DiffTypeDelete = DiffType("Delete")
	//DiffTypeChange shows some settings of this node is changed and existed before
	DiffTypeChange = DiffType("Change")
)

type PeerDiffrence struct {
	Diff DiffType
	NewPeer *spec.Peer
	OldPeer *spec.Peer
}

type Peers struct {
	mtx   sync.Mutex
	peers *spec.Peers
}

func NewPeers(peers *spec.Peers) *Peers {
	return &Peers{peers: peers}
}

type ClientStatus struct {
	HostName string
	IPAddr   net.IPNet
	ASN      *spec.ASN
	TunName  string
	Peers    *Peers
}

func NewClient(hostName string, tunName string, peers *spec.Peers) *ClientStatus {
	return &ClientStatus{HostName: hostName, TunName: tunName, Peers: NewPeers(peers)}
}

func (p *Peers) find(name string) (*spec.Peer, error) {
	for _, q := range p.peers.GetPeers() {
		if q.GetName() == name {
			return q, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (p *Peers) haveEqualName(q *spec.Peer) bool {
	for _, r := range p.peers.GetPeers() {
		if r.GetName() == q.GetName() {
			return true
		}
	}
	return false
}

func (p *Peers) haveEqual(q *spec.Peer) bool {
	for _, r := range p.peers.GetPeers() {
		if proto.Equal(r, q) {
			return true
		}
	}
	return false
}

func (p *Peers) createOrUpdateByName(q *spec.Peer) {
	for i, r := range p.peers.GetPeers() {
		if r.GetName() == q.GetName() {
			p.peers.Peers[i] = q
			return
		}
	}
	p.peers.Peers = append(p.peers.Peers, q)
}

func (c *ClientStatus) UpdatePeers(peers *Peers) ([]PeerDiffrence, error) {
	c.Peers.mtx.Lock()
	defer c.Peers.mtx.Unlock()
	for i, q := range peers.peers.GetPeers() {
		if q.Name == c.HostName {
			peers.peers.Peers = append(peers.peers.Peers[:i], peers.peers.Peers[i+1:]...)
			break
		}
	}
	ret := make([]PeerDiffrence, 0)
	//現在は存在するが次存在しないもの
	for i, q := range c.Peers.peers.GetPeers() {
		//現在と次のどちらにも存在するが名前以外の何かが変わったもの
		if peers.haveEqualName(q) && !peers.haveEqual(q) {
			ret = append(ret, PeerDiffrence{Diff: DiffTypeChange, NewPeer: peers.peers.GetPeers()[i], OldPeer: c.Peers.peers.Peers[i]})
		}
		if !peers.haveEqualName(q) && !peers.haveEqual(q) {
			ret = append(ret, PeerDiffrence{Diff: DiffTypeDelete, OldPeer: c.Peers.peers.GetPeers()[i]})
		}
	}
	//次存在するが現在存在しない
	for i, q := range peers.peers.GetPeers() {
		//changeは省く
		if !c.Peers.haveEqualName(q) && !c.Peers.haveEqual(q) {
			ret = append(ret, PeerDiffrence{Diff: DiffTypeAdd, NewPeer: peers.peers.GetPeers()[i]})
		}
	}
	c.Peers.peers = peers.peers
	return ret, nil
}

func (c *ClientStatus) GetPeers() *spec.Peers {
	c.Peers.mtx.Lock()
	defer c.Peers.mtx.Unlock()
	return c.Peers.peers
}
