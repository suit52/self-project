package main

import (
	"context"
	"fmt"
	"log"
	pb "self-project/calculator/pkg/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	var a, b int64
	fmt.Println("type values of a and b")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Sayadd(ctx, &pb.HelloRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Addition Result is : %v", r.GetC())

	r, err = c.Saymul(ctx, &pb.HelloRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("multiplication Result is : %v", r.GetC())
}
