package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gotti/meshover/internal/statusqueue"
	"github.com/gotti/meshover/spec"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	listenAddress     = flag.String("listen", "", "example: 0.0.0.0:8080")
	graphNodeConnNode = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "graph_meshover_nodes_nodes",
		Help: "Graph nodes of meshover nodes connection",
	},
		[]string{"id", "title", "mainstat"},
	)
	graphNodeConnEdge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "graph_meshover_nodes_bgp_edges",
		Help: "Graph edges of meshover nodes connection",
	},
		[]string{"id", "source", "target"},
	)
)

type server struct {
	spec.UnimplementedStatusManagerServiceServer
	queue *statusqueue.Queue
}

// RegisterStatus registers server status
func (s *server) RegisterStatus(ctx context.Context, req *spec.RegisterStatusRequest) (*spec.RegisterStatusResponse, error) {
	fmt.Println("received", req)
	s.queue.Add(req.GetStatus())
	return &spec.RegisterStatusResponse{}, nil
}

func (s *server) set() {
	for {
		prometheus.GaugeVec.Reset(*graphNodeConnEdge)
		prometheus.GaugeVec.Reset(*graphNodeConnNode)

		nodes, edges := s.queue.NodesAndEdges()
		fmt.Println("@@@@@@@@@@@@@@@@\nnodes")
		for _, n := range nodes {
			fmt.Println(n.Hostname)
			graphNodeConnNode.With(prometheus.Labels{"id": strconv.FormatUint(uint64(n.ID), 10), "title": n.Hostname, "mainstat": n.IP}).Set(0)
		}
		fmt.Println("edges")
		for _, e := range edges {
			fmt.Println(e.SourceID, e.TargetID)
			graphNodeConnEdge.With(prometheus.Labels{"id": e.ID, "source": strconv.FormatUint(uint64(e.SourceID), 10), "target": strconv.FormatUint(uint64(e.TargetID), 10)}).Set(0)
		}
		time.Sleep(15 * time.Second)
	}
}

func main() {
	flag.Parse()
	ctx := context.Background()
	server := server{queue: statusqueue.NewQueue(ctx, 30*time.Second)}
	s := grpc.NewServer()
	spec.RegisterStatusManagerServiceServer(s, &server)
	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	go server.set()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe("0.0.0.0:2112", nil); err != nil {
			log.Fatalln("failed to start http server")
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
