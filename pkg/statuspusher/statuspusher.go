package statuspusher

import (
	"context"
	"fmt"

	"github.com/gotti/meshover/spec"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	context   context.Context
	conn      *grpc.ClientConn
	grpcConn  spec.StatusManagerServiceClient
}

func NewClient(ctx context.Context, log *zap.Logger, statusServer string) (*Client, error) {
	conn, err := grpc.Dial(statusServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to statusServer, err=%w", err)
	}
	c := spec.NewStatusManagerServiceClient(conn)
	return &Client{context: ctx, conn: conn, grpcConn: c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) RegisterStatus(ctx context.Context, req *spec.RegisterStatusRequest) error{
	_, err := c.grpcConn.RegisterStatus(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to connect RegisterStatus endpont, err=%w", err)
	}
	return nil
}
