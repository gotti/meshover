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
	Use:   "gen",
	Short: "generate key",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial("localhost:12384", grpc.WithInsecure())
		if err != nil {
			log.Fatalln("failed to connect gRPC server", err)
		}
		defer cc.Close()
		c := spec.NewAdministratorServiceClient(cc)
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"Bearer": "aiueo"}))
		res, err := c.GenerateAgentKey(ctx, &spec.GenerateAgentKeyRequest{})
		if err != nil {
			log.Fatalln("failed to generate", err)
		}
		fmt.Printf("%x %x\n", res.GetAgentKey().GetId(), res.AgentKey.GetKey())
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
