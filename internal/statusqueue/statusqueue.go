package statusqueue

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gotti/meshover/spec"
)

type queueElement struct {
	hostname string
	status   *spec.StatusManagerPeerStatus
	ttl      time.Time
}

// Queue can store status and each element will discard after 1 minutes
// hostname is key and unique. if specified key is already exists, this will replace that
type Queue struct {
	mtx        sync.Mutex
	mapper     *IDMapper
	defaultttl time.Duration
	data       []*queueElement
}

// NewQueue creates queue, Added data will be deleted after ttl passed
func NewQueue(ctx context.Context, ttl time.Duration) *Queue {
	q := Queue{mtx: sync.Mutex{}, mapper: NewHostnameIDMapper(), defaultttl: ttl, data: []*queueElement{}}
	go q.gc(ctx)
	return &q
}

func (q *Queue) cleanUntil(until time.Time) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	for i, e := range q.data {
		if until.After(e.ttl) {
			q.data = append(q.data[:i], q.data[i+1:]...)
		}
	}
}

func (q *Queue) gc(ctx context.Context) {
	t := time.NewTicker(q.defaultttl / 5)
	defer func() {
		t.Stop()
	}()
	for {
		select {
		case <-t.C:
			now := time.Now()
			q.cleanUntil(now)
		case <-ctx.Done():
			break
		}
	}
}

// Add adds one element to queue
func (q *Queue) Add(d *spec.StatusManagerPeerStatus) error {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	for i, e := range q.data {
		if e.hostname == d.GetHostname() {
			q.data[i].status = d
			q.data[i].ttl = time.Now().Add(q.defaultttl)
			return nil
		}
	}
	q.data = append(q.data, &queueElement{hostname: d.GetHostname(), status: d, ttl: time.Now().Add(q.defaultttl)})
	return nil
}

// List shows Queue data list
func (q *Queue) List() []*spec.StatusManagerPeerStatus {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	var ret []*spec.StatusManagerPeerStatus
	for _, d := range q.data {
		ret = append(ret, d.status)
	}
	return ret
}

// Node is prometheus form node list
type Node struct {
	ID       uint32
	Hostname string
	IP       string
}

// Edge is prometheus form edge list
type Edge struct {
	ID       string
	SourceID uint32
	TargetID uint32
}

// NodesAndEdges converts received status to Nodes and Edges
func (q *Queue) NodesAndEdges() ([]*Node, []*Edge) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	n := q.nodes()
	e := q.edgesWithBGP()
	return n, e
}

func (q *Queue) nodes() []*Node {
	var ret []*Node
	for _, d := range q.data {
		ret = append(ret, &Node{ID: d.status.GetNodeStatus().LocalAS.GetNumber(), Hostname: d.status.GetHostname(), IP: d.status.GetNodeStatus().GetAddresses()[0].Format()})
	}
	return ret
}

func (q *Queue) edgesWithBGP() []*Edge {
	var ret []*Edge
	for _, d := range q.data {
		e := q.edgesOnOneNode(d.status)
		ret = append(ret, e...)
	}
	return ret
}

func (q *Queue) edgesOnOneNode(node *spec.StatusManagerPeerStatus) []*Edge {
	var ret []*Edge
	for _, d := range node.GetPeerStatus() {
		if d.GetBgpStatus().GetBGPState() == spec.BGPStates_BGPStateEstablished {
			ret = append(ret, &Edge{ID: q.mapper.Map(node.GetHostname() + "-" + d.GetRemoteHostname()), SourceID: node.GetNodeStatus().LocalAS.GetNumber(), TargetID: d.GetBgpStatus().RemoteAS.GetNumber()})
		}
	}
	return ret
}

// IDMapper maps hostname to unique integer id
type IDMapper struct {
	mtx    sync.Mutex
	mapper map[string]int
}

// NewHostnameIDMapper creates HostnameIDMapper
func NewHostnameIDMapper() *IDMapper {
	m := make(map[string]int)
	return &IDMapper{mtx: sync.Mutex{}, mapper: m}
}

// Map returns unique string id by hostname
func (m *IDMapper) Map(hostname string) string {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	id, ok := m.mapper[hostname]
	if !ok {
		i := m.getUniqueID()
		m.mapper[hostname] = i
		return strconv.Itoa(i)
	}
	return strconv.Itoa(id)
}

// RevMap returns hostname by unique string
func (q *Queue) RevMap(id string) string {
	idi, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid id", err)
	}
	q.mapper.mtx.Lock()
	defer q.mapper.mtx.Unlock()
	for k, v := range q.mapper.mapper {
		if v == idi {
			return k
		}
	}
	return "unknown"
}

// getUniqueID returns a unique id
func (m *IDMapper) getUniqueID() int {
	max := 0
	for _, id := range m.mapper {
		if id > max {
			max = id
		}
	}
	return max + 1
}
