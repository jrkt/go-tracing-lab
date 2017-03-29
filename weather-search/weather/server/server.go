package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"cloud.google.com/go/trace"
	auth "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/auth/client"
	cache "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/cache/client"
	http "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/http/client"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/interceptors"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/weather"
	pb "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/weather/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	port = "8002"
)

type server struct{}

func (s *server) GetCurrent(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {

	log.Println("requesting current weather data for zip:", in.Zip)
	time.Sleep(time.Duration(rand.Int31n(25)) * time.Millisecond)

	// authenticate request with provided token
	authClient, err := auth.New()
	if err != nil {
		log.Fatalf("failed to get weather client: %v", err)
	}
	authenticated, err := authClient.IsAuthenticated(ctx, in.Token)
	if err != nil || !authenticated {
		return nil, errors.New("failed to authenticate user: " + err.Error())
	}

	// check local cache for zip provided
	cacheClient, err := cache.New()
	if err != nil {
		log.Fatalf("failed to get cache client: %v", err)
	}
	data, err := cacheClient.Lookup(ctx, in.Zip)
	if err == nil || len(data) > 0 {
		// TODO: convert byte array to weather response
		return &pb.WeatherResponse{}, nil
	}

	// contact http service to perform request
	httpClient, err := http.New()
	if err != nil {
		log.Fatalf("failed to get http client: %v", err)
	}

	requestUrl := fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/q/UT/%d.json", in.Token, in.Zip)
	var conditions weather.Conditions
	err = httpClient.GET(ctx, requestUrl, &conditions)
	if err != nil {
		return nil, errors.New("http request failed: " + err.Error())
	}

	return &pb.WeatherResponse{
		Location:    conditions.CurrentObservation.DisplayLocation.Full,
		Description: conditions.CurrentObservation.Weather,
		Temperature: conditions.CurrentObservation.TempF,
	}, nil
}

func main() {

	// generate new tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// establish trace client
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY_FILE")))
	if err != nil {
		log.Fatalf("failed to establish trace client: %v", err)
	}

	// establish new gRPC server w/ custom server interceptor
	s := grpc.NewServer(interceptors.EnableGRPCTracingServerOption(tc))

	// register new server
	pb.RegisterWeatherServer(s, &server{})

	println("listening on :" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
