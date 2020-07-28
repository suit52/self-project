package main

import (
	"context"
	"log"
	"net"
	pb "self-project/calculator/pkg/proto"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Sayadd(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	a, b := in.GetA(), in.GetB()

	result := a + b
	log.Printf("Received: %v %v", a, b)
	return &pb.HelloReply{C: result}, nil
}
func (s *server) Saymul(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	a, b := in.GetA(), in.GetB()

	result := a * b
	log.Printf("Received: %v %v", a, b)
	return &pb.HelloReply{C: result}, nil
}
