package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prachin77/blog-web/db"
	"github.com/prachin77/blog-web/pb"
	"github.com/prachin77/blog-web/utils"
	"github.com/prachin77/blog-web/server/server_handlers"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// ✅ Initialize Firebase DBs
	db.Init()

	// ✅ Graceful cleanup on shutdown
	defer func() {
		log.Println("🔒 Closing database connections...")
		db.Close()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server_handlers.AuthServer{})
	pb.RegisterBlogServiceServer(s, &server_handlers.BlogServer{})

	log.Printf("✅ Server listening on port %d", config.ServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
