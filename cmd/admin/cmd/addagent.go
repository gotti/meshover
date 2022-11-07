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

var genCmd = &cobra.Command{
	Use:   "gen <ip to control server>, <admin token>",
	Short: "generate key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(args[0], grpc.WithInsecure())
		if err != nil {
			log.Fatalln("failed to connect gRPC server", err)
		}
		defer cc.Close()
		c := spec.NewAdministratorServiceClient(cc)
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"Bearer": args[1]}))
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
