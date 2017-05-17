package main

import (
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/trace"
	pb "github.com/jonathankentstevens/go-tracing-lab/grpc/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	headerKey = "stackdriver-trace-context"
	address   = "localhost:50051"
)

func main() {

	// establish connection with service w/ custom client interceptor
	conn, err := grpc.Dial(address, grpc.WithInsecure(), EnableGRPCTracingDialOption)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// register new connection as a GreeterClient
	c := pb.NewGreeterClient(conn)

	// establish new trace client
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
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

	// blocks until traces have been uploaded to GCP
	err = span.FinishWait() // use span.Finish() if your client is a long-running process.
	if err != nil {
		log.Fatalf("error finishing & uploading traces: %v", err)
	}

	println("Response:", r.Message)
}

// EnableGRPCTracingDialOption enables tracing of requests that are sent over a gRPC connection.
var EnableGRPCTracingDialOption = grpc.WithUnaryInterceptor(grpc.UnaryClientInterceptor(clientInterceptor))

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
	md[headerKey] = append(
		md[headerKey], fmt.Sprintf("%s/%d;o=1", span.TraceID(), 0),
	)
	ctx = metadata.NewContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
