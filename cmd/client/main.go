// Package main implements a client for Messenger service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/anthcor/go-grpc-template/messenger"

	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "We're on the air."
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessengerClient(conn)

	// Contact the server and print out its response.
	message := defaultMessage
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.MessageRequest{Message: message})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Message: %s", r.GetMessage())
}
