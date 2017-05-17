package main

import (
	"log"
	"net"
	"os"
	"strings"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	headerKey = "stackdriver-trace-context"
	port      = "50051"
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
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("failed to establish trace client: %v", err)
	}

	// establish new gRPC server w/ custom server interceptor
	s := grpc.NewServer(EnableGRPCTracingServerOption(tc))

	// register new server
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
		header := strings.Join(md[headerKey], "")

		// create new child span from google trace header & add to current request context
		span := traceClient.SpanFromHeader(info.FullMethod, header)
		defer span.Finish()

		// attach span to context from request
		ctx = trace.NewContext(ctx, span)

		return handler(ctx, req)
	}
}
