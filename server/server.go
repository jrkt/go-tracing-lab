package main

import (
	"log"
	"net"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	span := trace.FromContext(ctx)
	defer span.Finish()

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_SVCACCT_KEY")))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(trace.GRPCServerInterceptor(tc)))
	pb.RegisterGreeterServer(s, &server{})

	println("listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
