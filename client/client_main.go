package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"github.com/prachin77/blog-web/client/client_handlers"
	"github.com/prachin77/blog-web/pb"
	"github.com/prachin77/blog-web/utils"
	"github.com/prachin77/blog-web/middleware"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config: %v", err)
	}

	// Create gRPC connection
	grpcAddress := fmt.Sprintf("localhost:%d", config.ServerPort)
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("failed to connect: %v", err)
	}
	defer conn.Close()

	auth_service_client := pb.NewAuthServiceClient(conn)
	fmt.Println("AuthServiceClient created : ", auth_service_client)
	fmt.Printf("Connected to gRPC server on port %d", config.ServerPort)

	// 🔥 ADD THIS LINE - Initialize the client in handlers
	client_handlers.InitializeAuthClient(auth_service_client)

	// Create Gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	
	r.Use(middleware.RequestLogger())
	// r.Use(gin.Recovery())

	// Apply middlewares
	r.Use(middleware.CORSMiddleware()) 

	// Routes
	r.GET("/", client_handlers.DefaultRoute)
	r.GET("/blogger", client_handlers.RenderInitPage)
	r.GET("/auth", client_handlers.RenderAuthPage)

	// Auth form submission routes (HTMX targets)
	r.POST("/auth/register", client_handlers.Register)
	r.POST("/auth/login", client_handlers.Login)

	// 🔥 ADD DEBUG ROUTE (optional)
	r.GET("/debug/client", func(c *gin.Context) {
		if client_handlers.IsAuthClientInitialized() {
			c.JSON(200, gin.H{"client_status": "initialized and ready"})
		} else {
			c.JSON(200, gin.H{"client_status": "nil - not initialized"})
		}
	})

	fmt.Printf("🚀 Web server starting on port %d", config.ClientPort)
	fmt.Printf("🔗 Visit: http://localhost:%d", config.ClientPort)
	
	if err := r.Run(fmt.Sprintf(":%d", config.ClientPort)); err != nil {
		log.Fatalf("failed to start web server: %v", err)
	}
}