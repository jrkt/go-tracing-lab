package main

import (
	"log"
	"os"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/grpc-tracing-lab/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
)

func main() {

	// establish connection with service w/ custom client interceptor
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(trace.GRPCClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// register new connection as a GreeterClient
	c := pb.NewGreeterClient(conn)

	// establish new trace client
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("failed to establish new trace client: %v", err)
	}

	// create root span
	span := tc.NewSpan("/greeter/SayHello")
	span.SetLabel("from", "Erlich Bachman")

	// build span into context
	ctx = trace.NewContext(ctx, span)

	// pass new context into gRPC call to service
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Richard"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// blocks until traces have been uploaded to gcp
	err = span.FinishWait() // use span.Finish() if your client is a long-running process.
	if err != nil {
		log.Fatalf("error finishing & uploading traces: %v", err)
	}

	println("Response:", r.Message)
}
