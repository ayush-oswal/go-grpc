package main

import (
	"context"
	"log"

	"github.com/ayush-oswal/go-grpc/grpc-client/chat"
	"google.golang.org/grpc"
)

func sendRequest() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.NewClient("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance.
	client := chat.NewChatServiceClient(conn)

	// Prepare a message to send to the server.
	message := &chat.Message{
		Message: "Hello from client!",
	}

	// Call the SayHello RPC method on the server.
	response, err := client.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}

	// Log the server's response.
	log.Printf("Response from server: %s", response.Message)
}

func main() {
	sendRequest()
}
