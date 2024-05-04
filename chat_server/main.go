package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	chat "github.com/codingogo/go-grpc/chat/gen"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "gRPC server port")

type server struct {
	chat.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("received from client: %s", in.GetBody())
	return &chat.Message{Body: "Hello from server side!"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("faile to listen: %v", err)
	}

	s := grpc.NewServer()
	chat.RegisterGreeterServer(s, &server{})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
