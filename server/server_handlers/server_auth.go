package server_handlers

import (
	"context"

	"github.com/prachin77/blog-web/pb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Implement registration logic here
	return &pb.RegisterResponse{Message: "Registration successful"}, nil
}
