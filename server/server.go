package main

import (
	"log"
	"net"
	"os"
	"strings"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	println("HERE")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(EnableGRPCTracingServerOption(tc))
	pb.RegisterGreeterServer(s, &server{})

	println("listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// EnableGRPCTracingServerOption enables parsing google trace header from metadata
// and adds a new child span to the incoming request context.
func EnableGRPCTracingServerOption(traceClient *trace.Client) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(traceClient))
}

func serverInterceptor(traceClient *trace.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// fetch metadata from request context
		md, ok := metadata.FromContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		header := strings.Join(md["X-Cloud-Trace-Context"], "")

		// create new child span from google trace header, add to
		// current request context
		span := traceClient.SpanFromHeader(info.FullMethod, header)
		defer span.Finish()
		ctx = trace.NewContext(ctx, span)

		return handler(ctx, req)
	}
}
