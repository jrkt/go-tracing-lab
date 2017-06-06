package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/jonathankentstevens/go-tracing-lab/grpc/interceptors"
	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/weather-search/http/proto"
	"github.com/jonathankentstevens/go-tracing-lab/traceclient"
	"golang.org/x/net/context"
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
	tc, err := traceclient.New()
	if err != nil {
		log.Fatalln(err)
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
