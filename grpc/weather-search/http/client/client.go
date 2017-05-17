package client

import (
	"encoding/json"
	"errors"
	"time"

	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/weather-search/http/proto"
	"github.com/jonathankentstevens/go-tracing-lab/grpc/weather-search/interceptors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SvcClient struct {
	service pb.HTTPClient
}

var (
	address = "localhost:8004"
)

// New connects to the http service and initializes a client
func New() (*SvcClient, error) {

	timeout := grpc.WithTimeout(time.Second * 2)

	g, err := grpc.Dial(address, grpc.WithInsecure(), timeout, interceptors.EnableGRPCTracingDialOption)
	if err != nil {
		return nil, err
	}

	// get the service client
	c := SvcClient{
		service: pb.NewHTTPClient(g),
	}

	return &c, nil
}

// GET performs a HTTP GET request
func (c SvcClient) GET(ctx context.Context, url string, v interface{}) error {

	req := &pb.GetRequest{Url: url}

	res, err := c.service.GET(ctx, req)
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("no data returned")
	}

	return json.Unmarshal(res.Data, &v)
}
