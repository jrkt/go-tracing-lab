package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/grpc-tracing-lab/weather-search/auth/proto"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/interceptors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	port = "8001"
)

type server struct{}

// TODO: finish
func (s *server) IsAuthenticated(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("checking auth for token:", in.Token)
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return &pb.AuthResponse{Authenticated: true}, nil
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
	pb.RegisterAuthServer(s, &server{})

	println("listening on :" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
