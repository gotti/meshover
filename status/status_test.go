package status

import (
	"fmt"
	"net"
	"sync"
	"testing"

	"github.com/gotti/meshover/spec"
	"google.golang.org/protobuf/proto"
)


func isEqualPeerDiffArray(a,b []PeerDiffrence) error{
	if len(a) != len(b) {
		return fmt.Errorf("length of a and b not match")
	}
	for i,d := range a{
		if d.Add != b[i].Add{
			return fmt.Errorf("in array of diff, add not match")
		}
		if !proto.Equal(d.Peer, b[i].Peer){
			return fmt.Errorf("in array of diff, %+v and %+v not matched", d.Peer, b[i].Peer)
		}
	}
	return nil
}

func TestUpdatePeers(t *testing.T){
	peerSelf := &spec.Peer{
				Name: "host",
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
				Name: "host1",
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
				Name: "host2",
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
	data := []struct{
		testName string
		initialPeers *Peers
		nextPeers *Peers
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
					Add: true,
					Peer: peerB,
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
					Add: false,
					Peer: peerB,
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
					Peers: []*spec.Peer{
					},
				},
			},
			expectedDiff: []PeerDiffrence{
			},
		},
	}
	for _, d := range data{
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
		}
	}
}
