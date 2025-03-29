package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/grpc/proto" // Import generated protobuf package
	"google.golang.org/grpc"
)

// Implement the Greeter service
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Received request from:", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	// Start a TCP listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	fmt.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
