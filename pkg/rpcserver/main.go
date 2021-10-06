package main

import (
	"context"
	"log"
	"net"

	"github.com/diwangtseb/airchat/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type commonServer struct {
	pb.UnimplementedCommonServer
}

func (commonServer) Init(context.Context, *pb.Empty) (*pb.Empty, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommonServer(s, &commonServer{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
