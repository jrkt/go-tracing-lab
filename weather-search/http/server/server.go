package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/http/proto"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/interceptors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	port = "8004"
)

type server struct{}

// TODO: finish
func (s *server) GET(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {

	// make HTTP GET request for current weather conditions
	log.Println("request url:", in.Url)
	req, err := http.Get(in.Url)
	if err != nil {
		return nil, errors.New("http GET failed: " + err.Error())
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("invalid response body: " + err.Error())
	}
	req.Body.Close()

	return &pb.GetResponse{Data: b}, nil
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
	pb.RegisterHTTPServer(s, &server{})

	println("listening on :" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
