package main

import (
	"log"

	pb "github.com/jonathankentstevens/grpc-tracing-lab/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
)

func main() {

	// establish connection with service w/ custom client interceptor
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// register new connection as a GreeterClient
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	// call service
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Richard"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	println("Response:", r.Message)
}
