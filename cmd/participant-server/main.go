package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/TastyPi/grail-interview/internal/participant/server"
)

const (
	port = ":50051"
)

func main() {
	// Start listening for requests.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v", port)
	}
	log.Println("Listening on port", port)

	// Register the service.
	s := grpc.NewServer()
	pb.RegisterParticipantServiceServer(s, server.Create())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
