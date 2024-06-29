package main

import (
	"context"
	"log"
	"net"

	"github.com/ayush-oswal/go-grpc/grpc-server/chat"
	"google.golang.org/grpc"
)

type ChatServer struct {
	chat.UnimplementedChatServiceServer
}

func (s ChatServer) SayHello(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	log.Printf("Recieved Message from Client: %v", message.Message)
	return &chat.Message{Message: "Hello World"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Server cannot listen on port 9000")
	}

	grpcServer := grpc.NewServer()

	service := &ChatServer{}

	chat.RegisterChatServiceServer(grpcServer, service)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("grpcServer cannot be serverd! : %s", err)
	}

}
