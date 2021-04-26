// Package main implements a server for Messenger service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/anthcor/go-grpc-template/messenger"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement messenger.MessengerServer.
type server struct {
	pb.UnimplementedMessengerServer
}

// SendServerMessage implements messenger.MessengerServer
func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.MessageResponse{Message: "The server received your message: " + in.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
