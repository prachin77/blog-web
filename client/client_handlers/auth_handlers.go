package client_handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/prachin77/blog-web/model"
    "github.com/prachin77/blog-web/pb"
)

var (
    auth_service_client pb.AuthServiceClient
)

const CONTEXT_TIMEOUT = 5 * time.Second

type RegisterResponse struct {
    Message string `json:"message"`
    UserId  string `json:"user_id"`
}

// InitializeAuthClient initializes the gRPC client
func InitializeAuthClient(client pb.AuthServiceClient) {
    auth_service_client = client
    // Add a log to confirm initialization
    if client != nil {
        println("✅ AuthServiceClient initialized successfully")
    } else {
        println("❌ AuthServiceClient initialization failed - client is nil")
    }
}

// IsAuthClientInitialized checks if the auth client is initialized
func IsAuthClientInitialized() bool {
    return auth_service_client != nil
}

func Register(ctx *gin.Context) {
    // Check if client is initialized
    if auth_service_client == nil {
        println("❌ Register called but auth_service_client is nil")
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Auth service client not initialized",
            "hint": "Make sure the gRPC server is running and client is properly initialized",
        })
        return
    }

    println("✅ Register called with initialized client")
    ctx.Header("Content-Type", "application/json")

    var user model.User
    err := json.NewDecoder(ctx.Request.Body).Decode(&user)
    if err != nil {
        println("❌ JSON decode error:", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    println("📝 Registration request for:", user.Username, user.Email)

    if user.Username == "" || user.Password == "" || user.Email == "" {
        println("❌ Validation failed - empty fields")
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username, password or email is empty"})
        return
    }

    req := &pb.RegisterRequest{
        Username:  user.Username,
        Password:  user.Password,
        Email:     user.Email,
        CreatedAt: time.Now().Format(time.RFC3339),
        NoOfBlogs: 0,
        Followers: 0,
    }

    println("🔄 Calling gRPC Register service...")

    ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
    defer cancelFunc()

    res, err := auth_service_client.Register(ctxTimeout, req)
    if err != nil {
        println("❌ gRPC call failed:", err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to register user", 
            "details": err.Error(),
        })
        return
    }

    println("✅ User registration successful, Firebase ID:", res.UserId)

    ctx.JSON(http.StatusOK, RegisterResponse{
        Message: res.Message,
        UserId:  res.UserId,
    })
}