package client

import (
	"time"

	"github.com/jrkt/go-tracing-lab/grpc/interceptors"
	pb "github.com/jrkt/go-tracing-lab/grpc/weather-search/auth/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SvcClient struct {
	service pb.AuthClient
}

var (
	address = "localhost:8001"
)

// New connects to the auth service and initializes a client
func New() (*SvcClient, error) {

	timeout := grpc.WithTimeout(time.Second * 2)

	g, err := grpc.Dial(address, grpc.WithInsecure(), timeout, interceptors.EnableGRPCTracingDialOption)
	if err != nil {
		return nil, err
	}

	// get the service client
	c := SvcClient{
		service: pb.NewAuthClient(g),
	}

	return &c, nil
}

// IsAuthenticated calls the auth service with the token provided to check authentication
func (c SvcClient) IsAuthenticated(ctx context.Context, token string) (bool, error) {

	req := &pb.AuthRequest{Token: token}

	res, err := c.service.IsAuthenticated(ctx, req)
	if err != nil {
		return false, err
	}

	return res.Authenticated, nil
}
