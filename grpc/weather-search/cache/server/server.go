package main

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/weather-search/cache/proto"
	"github.com/jonathankentstevens/go-tracing-lab/grpc/interceptors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
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
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("failed to establish trace client: %v", err)
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
