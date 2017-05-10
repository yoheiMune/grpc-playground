package main

import (
	pb "github.com/yoheiMune/grpc-playground/helloworld"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	// Contact the server and print its response.
	name := "Yohei"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.Message)

}