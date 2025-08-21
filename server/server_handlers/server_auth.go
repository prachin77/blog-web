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
    emailExists,user_id ,err := db.CheckEmailExists(req.Email)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Error checking email existence: %v", err)
    }

    if emailExists {
        return &pb.RegisterResponse{
            Message: "User with this email already exists",
            UserId:  user_id, 
        }, nil
    }

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

    userID, err := db.CreateUser(user)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
    }

    return &pb.RegisterResponse{
        Message: "User registered successfully",
        UserId:  userID, 
    }, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    // check through email
    // if email exists then user can log in 

    emailExists, userID, err := db.CheckEmailExists(req.Email) 
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Error checking email existence: %v", err)
    }

    if !emailExists {
        return nil, status.Errorf(codes.NotFound, "User with this email does not exist")
    }

    passwordMatch, err := db.VerifyPassword(userID, req.Password)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Error verifying password: %v", err)
    }

    if !passwordMatch {
        return nil, status.Errorf(codes.Unauthenticated, "Invalid password")
    }
    
    return &pb.LoginResponse{
        Message: "User successfully logged in",
        UserId:  userID, 
    }, nil
}
