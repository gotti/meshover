package status

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"

	"github.com/gotti/meshover/spec"
	"google.golang.org/protobuf/proto"
)

func isEqualPeerDiffArray(a, b []PeerDiffrence) error {
	if len(a) != len(b) {
		return fmt.Errorf("length of a and b not match")
	}
	for i, d := range a {
		if d.Diff != b[i].Diff {
			return fmt.Errorf("in array of diff, add not match")
		}
		if !proto.Equal(d.NewPeer, b[i].NewPeer) {
			return fmt.Errorf("NewPeer, in array of diff, %+v and %+v not matched", d.NewPeer, b[i].NewPeer)
		}
		if !proto.Equal(d.OldPeer, b[i].OldPeer) {
			return fmt.Errorf("OldPeer, in array of diff, %+v and %+v not matched", d.NewPeer, b[i].NewPeer)
		}
	}
	return nil
}

func TestUpdatePeers(t *testing.T) {
	peerSelf := &spec.Peer{
		Name:     "host",
		Asnumber: &spec.ASN{Number: 1},
		Address: []*spec.AddressCIDR{
			{
				Addresscidr: &spec.AddressCIDR_AddressCIDRIPv4{
					AddressCIDRIPv4: &spec.AddressCIDRIPv4{Ipaddress: &spec.AddressIPv4{
						Ipaddress: "10.1.0.4/32",
					}},
				},
			},
		},
		Underlay: &spec.Peer_UnderlayLinuxKernelWireguard{
			UnderlayLinuxKernelWireguard: &spec.UnderlayLinuxKernelWireguard{
				Endpoint: &spec.AddressAndPort{
					Ipaddress: spec.NewAddress(net.ParseIP("192.168.0.4")),
				},
			},
		},
	}

	peerA := &spec.Peer{
		Name:     "host1",
		Asnumber: &spec.ASN{Number: 1},
		Address: []*spec.AddressCIDR{
			{
				Addresscidr: &spec.AddressCIDR_AddressCIDRIPv4{
					AddressCIDRIPv4: &spec.AddressCIDRIPv4{Ipaddress: &spec.AddressIPv4{
						Ipaddress: "10.1.0.5/32",
					}},
				},
			},
		},
		Underlay: &spec.Peer_UnderlayLinuxKernelWireguard{
			UnderlayLinuxKernelWireguard: &spec.UnderlayLinuxKernelWireguard{
				Endpoint: &spec.AddressAndPort{
					Ipaddress: spec.NewAddress(net.ParseIP("192.168.0.5")),
				},
			},
		},
	}
	peerB := &spec.Peer{
		Name:     "host2",
		Asnumber: &spec.ASN{Number: 1},
		Address: []*spec.AddressCIDR{
			{
				Addresscidr: &spec.AddressCIDR_AddressCIDRIPv4{
					AddressCIDRIPv4: &spec.AddressCIDRIPv4{Ipaddress: &spec.AddressIPv4{
						Ipaddress: "10.1.0.6/32",
					}},
				},
			},
		},
		Underlay: &spec.Peer_UnderlayLinuxKernelWireguard{
			UnderlayLinuxKernelWireguard: &spec.UnderlayLinuxKernelWireguard{
				Endpoint: &spec.AddressAndPort{
					Ipaddress: spec.NewAddress(net.ParseIP("192.168.0.6")),
				},
			},
		},
	}
	peerB2 := &spec.Peer{
		Name:     "host2",
		Asnumber: &spec.ASN{Number: 1},
		Address: []*spec.AddressCIDR{
			{
				Addresscidr: &spec.AddressCIDR_AddressCIDRIPv4{
					AddressCIDRIPv4: &spec.AddressCIDRIPv4{Ipaddress: &spec.AddressIPv4{
						Ipaddress: "10.1.0.6/32",
					}},
				},
			},
		},
		Underlay: &spec.Peer_UnderlayLinuxKernelWireguard{
			UnderlayLinuxKernelWireguard: &spec.UnderlayLinuxKernelWireguard{
				Endpoint: &spec.AddressAndPort{
					Ipaddress: spec.NewAddress(net.ParseIP("192.168.0.6")),
				},
				PublicKey: &spec.Curve25519Key{Key: []byte("deadbeef")},
			},
		},
	}
	data := []struct {
		testName     string
		initialPeers *Peers
		nextPeers    *Peers
		expectedDiff []PeerDiffrence
	}{
		{
			testName: "no change",
			initialPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
					},
				},
			},
			nextPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
					},
				},
			},
			expectedDiff: []PeerDiffrence{},
		},
		{
			testName: "add peerB",
			initialPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
					},
				},
			},
			nextPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
						peerB,
					},
				},
			},
			expectedDiff: []PeerDiffrence{
				{
					Diff: DiffTypeAdd,
					NewPeer: peerB,
				},
			},
		},
		{
			testName: "del peerB",
			initialPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
						peerB,
					},
				},
			},
			nextPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
					},
				},
			},
			expectedDiff: []PeerDiffrence{
				{
					Diff: DiffTypeDelete,
					OldPeer: peerB,
				},
			},
		},
		{
			testName: "no change with self",
			initialPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerSelf,
					},
				},
			},
			nextPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{},
				},
			},
			expectedDiff: []PeerDiffrence{},
		},
		{
			testName: "change to b2",
			initialPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
						peerB,
					},
				},
			},
			nextPeers: &Peers{
				mtx: sync.Mutex{},
				peers: &spec.Peers{
					Peers: []*spec.Peer{
						peerA,
						peerB2,
					},
				},
			},
			expectedDiff: []PeerDiffrence{
				{
					Diff: DiffTypeChange,
					OldPeer: peerB,
					NewPeer: peerB2,
				},
			},
		},
	}
	for _, d := range data {
		c := NewClient("host", "meshover0", nil)
		_, err := c.UpdatePeers(d.initialPeers)
		if err != nil {
			t.Errorf("%s: an error occured, err=%s", d.testName, err)
		}
		diff, err := c.UpdatePeers(d.nextPeers)
		if err != nil {
			t.Errorf("%s: an error occured, err=%s", d.testName, err)
		}
		if err := isEqualPeerDiffArray(diff, d.expectedDiff); err != nil {
			t.Errorf("%s: got=%+v, expected=%+v, err=%s", d.testName, diff, d.expectedDiff, err)
			log.Fatalf("%s: got=%+v, expected=%+v, err=%s", d.testName, diff, d.expectedDiff, err)
		}
	}
}
