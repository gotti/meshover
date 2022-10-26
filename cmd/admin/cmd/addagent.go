package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/gotti/meshover/spec"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const hostp = "home.meshover.rax.rip:80"

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate key",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(hostp, grpc.WithInsecure())
		if err != nil {
			log.Fatalln("failed to connect gRPC server", err)
		}
		defer cc.Close()
		c := spec.NewAdministratorServiceClient(cc)
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"Bearer": "d283ddeb17748d6dee97e20c486f04ab05a4d3056d3954f1bc0de78ce0cf2c57"}))
		res, err := c.GenerateAgentKey(ctx, &spec.GenerateAgentKeyRequest{})
		if err != nil {
			log.Fatalln("failed to generate", err)
		}
		fmt.Printf("%x %s\n", res.GetAgentKey().GetId(), res.AgentKey.GetKey())
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
