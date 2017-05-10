package main

import (
	pb "github.com/yoheiMune/grpc-playground/helloworld"
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement hello.HelloServer.
type server struct {}

// SayHello implements hello.HelloServer.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen; %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

