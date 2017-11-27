package client

import (
	"time"

	"github.com/jrkt/go-tracing-lab/grpc/interceptors"
	pb "github.com/jrkt/go-tracing-lab/grpc/weather-search/weather/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SvcClient struct {
	service pb.WeatherClient
}

var (
	address = "localhost:8002"
)

// New connects to the weather service and initializes a client
func New() (*SvcClient, error) {

	timeout := grpc.WithTimeout(time.Second * 2)

	g, err := grpc.Dial(address, grpc.WithInsecure(), timeout, interceptors.EnableGRPCTracingDialOption)
	if err != nil {
		return nil, err
	}

	// get the service client
	c := SvcClient{
		service: pb.NewWeatherClient(g),
	}

	return &c, nil
}

func (c SvcClient) SearchByZip(ctx context.Context, token string, zip int64) (*pb.WeatherResponse, error) {
	req := &pb.WeatherRequest{Token: token, Zip: zip}

	return c.service.GetCurrent(ctx, req)
}
