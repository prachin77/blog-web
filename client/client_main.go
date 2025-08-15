// package main

// import (
// 	"fmt"
// 	"log"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	"github.com/gin-gonic/gin"
// 	"github.com/prachin77/blog-web/client/client_handlers"
// 	"github.com/prachin77/blog-web/pb"
// 	"github.com/prachin77/blog-web/utils"
// 	"github.com/prachin77/blog-web/middleware"
// )

// func main() {
// 	config, err := utils.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("Failed to load config: %v", err)
// 	}

// 	// Create gRPC connection
// 	grpcAddress := fmt.Sprintf("localhost:%d", config.ServerPort)
// 	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("failed to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	auth_service_client := pb.NewAuthServiceClient(conn)
// 	fmt.Println("AuthServiceClient created : ", auth_service_client)
// 	log.Printf("Connected to gRPC server on port %d", config.ServerPort)

// 	// Create Gin server
// 	gin.SetMode(gin.ReleaseMode)
//     r := gin.New()
//     r.Use(gin.Recovery())

// 	// Apply middlewares
// 	r.Use(middleware.CORSMiddleware())           // Enable CORS
// 	// r.Use(middleware.SecureMiddleware())         // Add security headers
//     // r.Use(middleware.RequestLoggingMiddleware()) // Log requests

//     r.GET("/", client_handlers.DefaultRoute)
//     r.GET("/blogger", client_handlers.RenderInitPage)
//     r.GET("/auth", client_handlers.RenderAuthPage)

//     // Auth form submission routes (HTMX targets)
//     r.POST("/auth/login", client_handlers.PostLogin)
//     r.POST("/auth/register", client_handlers.PostRegister)

// 	log.Printf("Web server starting on port %d", config.ClientPort)
// 	if err := r.Run(fmt.Sprintf(":%d", config.ClientPort)); err != nil {
// 		log.Fatalf("failed to start web server: %v", err)
// 	}
// }


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
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create gRPC connection
	grpcAddress := fmt.Sprintf("localhost:%d", config.ServerPort)
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	auth_service_client := pb.NewAuthServiceClient(conn)
	fmt.Println("AuthServiceClient created : ", auth_service_client)
	log.Printf("Connected to gRPC server on port %d", config.ServerPort)

	// Create Gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	
	// Add request logging middleware to debug HTMX requests
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	
	r.Use(gin.Recovery())

	// Apply middlewares
	r.Use(middleware.CORSMiddleware()) // Enable CORS

	// Routes
	r.GET("/", client_handlers.DefaultRoute)
	r.GET("/blogger", client_handlers.RenderInitPage)
	r.GET("/auth", client_handlers.RenderAuthPage)

	// Auth form submission routes (HTMX targets)
	r.POST("/auth/login", client_handlers.PostLogin)
	r.POST("/auth/register", client_handlers.PostRegister)

	log.Printf("🚀 Web server starting on port %d", config.ClientPort)
	log.Printf("🔗 Visit: http://localhost:%d", config.ClientPort)
	
	if err := r.Run(fmt.Sprintf(":%d", config.ClientPort)); err != nil {
		log.Fatalf("failed to start web server: %v", err)
	}
}