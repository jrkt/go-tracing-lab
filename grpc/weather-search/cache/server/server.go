package main

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/jrkt/go-tracing-lab/grpc/interceptors"
	pb "github.com/jrkt/go-tracing-lab/grpc/weather-search/cache/proto"
	"github.com/jrkt/go-tracing-lab/traceclient"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = "8003"
)

type server struct{}

// TODO: finish
func (s *server) Lookup(ctx context.Context, in *pb.LookupRequest) (*pb.LookupResponse, error) {
	log.Println("checking local cache for key:", in.Key)
	time.Sleep(time.Duration(rand.Int31n(35)) * time.Millisecond)
	return &pb.LookupResponse{}, errors.New("not found")
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
	pb.RegisterCacheServer(s, &server{})

	println("listening on :" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
