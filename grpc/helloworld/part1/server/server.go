package main

import (
	"log"
	"net"
	"os"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = "50051"
)

// server is used to implement helloworld.GreeterServer
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// SayHelloAgain implements helloworld.GreeterServer
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {

	// generate new tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// establish trace client
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("failed to establish trace client: %v", err)
	}

	// establish new gRPC server w/ custom server interceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(trace.GRPCServerInterceptor(tc)))

	// register new server
	pb.RegisterGreeterServer(s, &server{})

	println("listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
