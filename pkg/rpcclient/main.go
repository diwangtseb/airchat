package main

import (
	"context"
	"github.com/diwangtseb/airchat/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
)
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCommonClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Init(ctx,nil)
	log.Println(r)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}