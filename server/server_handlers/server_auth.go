package server_handlers

import (
    "context"
    "time"

    "github.com/prachin77/blog-web/db"
    "github.com/prachin77/blog-web/model"
    "github.com/prachin77/blog-web/pb"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type AuthServer struct {
    pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    // Check if email already exists
    emailExists, err := db.CheckEmailExists(req.Email)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Error checking email existence: %v", err)
    }

    if emailExists {
        return &pb.RegisterResponse{
            Message: "User with this email already exists",
            UserId:  "", // Empty string for existing user
        }, nil
    }

    // Parse the created_at time
    createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid created_at format: %v", err)
    }

    // Create user object (without ID - Firebase will generate it)
    user := &model.User{
        Username:  req.Username,
        Email:     req.Email,
        Password:  req.Password, // In production, hash this password before storing
        CreatedAt: createdAt,
        NoOfBlogs: int(req.NoOfBlogs), // Fixed: was using Followers instead of NoOfBlogs
        Followers: int(req.Followers),
    }

    // Save user to database and get Firebase-generated ID
    userID, err := db.CreateUser(user)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
    }

    return &pb.RegisterResponse{
        Message: "User registered successfully",
        UserId:  userID, // Return the actual Firebase document ID
    }, nil
}
