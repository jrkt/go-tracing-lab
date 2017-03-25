package main

import (
	"fmt"
	pb "grpc/helloworld/helloworld"
	"log"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc.UnaryClientInterceptor(clientInterceptor)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	tc, err := trace.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	span := tc.NewSpan("/greeter/SayHello")
	defer span.Finish()

	ctx = trace.NewContext(ctx, span)

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	println("Response:", r.Message)
}

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// trace current request w/ child span
	span := trace.FromContext(ctx).NewChild(method)
	defer span.Finish()

	// new metadata, or copy of existing
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	// append trace header to context metadata
	// header specification: https://cloud.google.com/trace/docs/faq
	md["X-Cloud-Trace-Context"] = append(
		md["X-Cloud-Trace-Context"], fmt.Sprintf("%s/%d;o=1", span.TraceID(), 0),
	)
	ctx = metadata.NewContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
