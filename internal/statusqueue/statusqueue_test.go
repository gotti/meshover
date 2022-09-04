package statusqueue

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/gotti/meshover/spec"
	"google.golang.org/protobuf/proto"
)

func TestListQueue(t *testing.T){
	ctx := context.Background()
	queue := NewQueue(ctx, 100*time.Millisecond)
	d := &spec.StatusManagerPeerStatus{
		Hostname: "hoge1",
	}
	queue.Add(d)
	l1 := queue.List()
	if !proto.Equal(d,l1[0]){
		t.Errorf("queue data not match expected %+v, got %+v", d, l1[0])
	}
	time.Sleep(150*time.Millisecond)
	l2 := queue.List()
	if len(l2) != 0{
		t.Errorf("queue data not match expected cleaned, got %+v", l2[0])
	}
}

func equalNodeArray(a, b []*Node) error{
	if len(a) != len(b){
		return fmt.Errorf("length not match, left=%d, right=%d", len(a), len(b))
	}
	for i := range a{
		if !reflect.DeepEqual(a[i], b[i]){
			return fmt.Errorf("element not match, index=%d left=%+v, right=%+v", i, a[i], b[i])
		}
	}
	return nil
}

func equalEdgeArray(a, b []*Edge) error{
	if len(a) != len(b){
		return fmt.Errorf("length not match, left=%d, right=%d", len(a), len(b))
	}
	for i := range a{
		if !reflect.DeepEqual(a[i], b[i]){
			return fmt.Errorf("element not match, index=%d left=%+v, right=%+v", i, a[i], b[i])
		}
	}
	return nil
}

func TestNodeList(t *testing.T){
	_, n1, _ := net.ParseCIDR("10.34.81.156/32") 
	data := []struct{
		testName string
		in []*spec.StatusManagerPeerStatus
		out []*Node
	}{
		{
			testName: "1",
			in: []*spec.StatusManagerPeerStatus{
				{
					Hostname: "hoge1",
					NodeStatus: &spec.MinimumNodeStatus{
						LocalAS: &spec.ASN{Number: 1},
						Addresses: []*spec.AddressCIDR{
							spec.NewAddressCIDR(*n1),
						},
					},
				},
			},
			out: []*Node{
				{
					ID: "1",
					Hostname: "hoge1",
					IP: "10.34.81.156/32",
				},
			},
		},
	}
	ctx := context.Background()
	for _, d := range data{
		queue := NewQueue(ctx, 999*time.Second)
		for _, n := range d.in{
			queue.Add(n)
		}
		if err := equalNodeArray(d.out, queue.nodes()); err != nil{
			t.Errorf("%s: err=%s", d.testName, err)
		}
	}
}

var (
	rawnet1 = "10.34.81.156/32"
	_, net1, _ = net.ParseCIDR(rawnet1)
	rawnet2 = "10.34.81.157/32"
	_, net2, _ = net.ParseCIDR(rawnet2)
	rawnet3 = "10.34.81.157/32"
	_, net3, _ = net.ParseCIDR(rawnet3)
	hosthoge1 = &spec.StatusManagerPeerStatus{
					Hostname: "hoge1",
					NodeStatus: &spec.MinimumNodeStatus{
						LocalAS: &spec.ASN{Number: 1},
						Addresses: []*spec.AddressCIDR{
							spec.NewAddressCIDR(*net1),
						},
					},
					PeerStatus: []*spec.NodePeersStatus{
						{
							RemoteHostname: "hoge2",
							BgpStatus: &spec.PeerBGPStatus{
								BGPState: spec.BGPStates_BGPStateEstablished,
							},
						},
					},
				}
	hosthoge2 = &spec.StatusManagerPeerStatus{
					Hostname: "hoge2",
					NodeStatus: &spec.MinimumNodeStatus{
						LocalAS: &spec.ASN{Number: 2},
						Addresses: []*spec.AddressCIDR{
							spec.NewAddressCIDR(*net2),
						},
					},
					PeerStatus: []*spec.NodePeersStatus{
						{
							RemoteHostname: "hoge1",
							BgpStatus: &spec.PeerBGPStatus{
								BGPState: spec.BGPStates_BGPStateEstablished,
							},
						},
					},
				}
)

func TestEdgesOnOneNode(t *testing.T){
	data := []struct{
		testName string
		in *spec.StatusManagerPeerStatus
		out []*Edge
	}{
		{
			testName: "1",
			in: hosthoge1,
			out: []*Edge{
				{
					ID: "1",
					SourceID: "2",
					TargetID: "3",
				},
			},
		},
	}
	ctx := context.Background()
	for _, d := range data{
		queue := NewQueue(ctx, 999*time.Second)
		e := queue.edgesOnOneNode(d.in)
		if err := equalEdgeArray(e, d.out); err != nil{
			t.Errorf("%s: err=%s", d.testName, err)
		}
	}

}

func TestNodeAndEdgeList(t *testing.T){
	data := []struct{
		testName string
		in []*spec.StatusManagerPeerStatus
		outNodes []*Node
		outEdges []*Edge
	}{
		{
			testName: "1",
			in: []*spec.StatusManagerPeerStatus{
				hosthoge1,
				hosthoge2,
				{
					Hostname: "hoge3",
					NodeStatus: &spec.MinimumNodeStatus{
						LocalAS: &spec.ASN{Number: 3},
						Addresses: []*spec.AddressCIDR{
							spec.NewAddressCIDR(*net3),
						},
					},
				},
			},
			outNodes: []*Node{
				{
					ID: "1",
					Hostname: "hoge1",
					IP: rawnet1,
				},
				{
					ID: "2",
					Hostname: "hoge2",
					IP: rawnet2,
				},
				{
					ID: "3",
					Hostname: "hoge3",
					IP: rawnet3,
				},
			},
			outEdges: []*Edge{
				{
					ID: "4",
					SourceID: "1",
					TargetID: "2",
				},
				{
					ID: "5",
					SourceID: "2",
					TargetID: "1",
				},
			},
		},
	}
	ctx := context.Background()
	for _, d := range data{
		queue := NewQueue(ctx, 999*time.Second)
		for _, n := range d.in{
			queue.Add(n)
		}
		nodes, edges := queue.NodesAndEdges()
		if err := equalNodeArray(nodes, d.outNodes); err != nil {
			t.Errorf("%s: err=%s", d.testName, err)
		}
		if err := equalEdgeArray(edges, d.outEdges); err != nil {
			t.Errorf("%s: err=%s", d.testName, err)
		}
	}
}
