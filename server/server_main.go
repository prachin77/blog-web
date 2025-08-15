package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prachin77/blog-web/pb"
	"github.com/prachin77/blog-web/server/server_handlers"
	"github.com/prachin77/blog-web/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server_handlers.AuthServer{})

	log.Printf("Server listening on port %d", config.ServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}