package interceptors

import (
	"fmt"
	"strings"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	headerKey = "stackdriver-trace-context"

	// EnableGRPCTracingDialOption enables tracing of requests that are sent over a gRPC connection.
	EnableGRPCTracingDialOption = grpc.WithUnaryInterceptor(grpc.UnaryClientInterceptor(clientInterceptor))
)

/*
// establish connection with service w/ custom client interceptor
conn, err := grpc.Dial(address, grpc.WithInsecure(), EnableGRPCTracingDialOption)
...
// initialize trace client
ctx := context.Background()
tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
if err != nil {
	log.Fatalf("failed to establish new trace client: %v", err)
}

// create root span
span := tc.NewSpan("/greeter/SayHello")
span.SetLabel("from", "Erlich Bachman")

// build span into context that is passed in gRPC request
ctx = trace.NewContext(ctx, span)
...
// make gRPC request
...
// blocks until traces have been uploaded to GCP
err = span.FinishWait() // use span.Finish() if your client is a long-running process.
*/
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

/*
// initialize trace client
ctx := context.Background()
tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
if err != nil {
	log.Fatalf("failed to establish trace client: %s", err)
}

// establish new gRPC server w/ custom server interceptor
s := grpc.NewServer(EnableGRPCTracingServerOption(tc))
*/

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
