package main

import (
	pb "grpc_tracer/proto"
	"log"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(trace.GRPCClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_SVCACCT_KEY")))
	if err != nil {
		log.Fatalln(err)
	}

	span := tc.NewSpan("/greeter/SayHello")

	ctx = trace.NewContext(ctx, span)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	err = span.FinishWait()
	if err != nil {
		log.Fatalln(err)
	}

	println("Response:", r.Message)
}
