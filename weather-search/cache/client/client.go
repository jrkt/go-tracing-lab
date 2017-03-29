package client

import (
	"time"

	pb "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/cache/proto"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/interceptors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SvcClient struct {
	service pb.CacheClient
}

var (
	address = "localhost:8003"
)

// New connects to the cache service and initializes a client
func New() (*SvcClient, error) {

	timeout := grpc.WithTimeout(time.Second * 2)

	g, err := grpc.Dial(address, grpc.WithInsecure(), timeout, interceptors.EnableGRPCTracingDialOption)
	if err != nil {
		return nil, err
	}

	// get the service client
	c := SvcClient{
		service: pb.NewCacheClient(g),
	}

	return &c, nil
}

// Lookup calls the cache service to find the associated value from the provided key
func (c SvcClient) Lookup(ctx context.Context, key int64) ([]byte, error) {

	req := &pb.LookupRequest{Key: key}
	res, err := c.service.Lookup(ctx, req)
	if err != nil {
		return nil, err
	}

	var b []byte
	if res != nil {
		b = res.Val
	}

	return b, err
}
