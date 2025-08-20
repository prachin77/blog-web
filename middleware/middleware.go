package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Collect log data
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		latency := time.Since(start)
		clientIP := c.ClientIP()

		// Skip logging for certain paths
		if path == "/health" || path == "/favicon.ico" {
			return
		}

		// Create colored boxes based on request type
		var methodBox, statusBox string

		switch method {
		case "GET":
			methodBox = color.New(color.BgBlue, color.FgWhite).Sprintf(" GET ")
		case "POST":
			methodBox = color.New(color.BgGreen, color.FgWhite).Sprintf(" POST ")
		case "PUT", "PATCH":
			methodBox = color.New(color.BgYellow, color.FgBlack).Sprintf(" %s ", method)
		case "DELETE":
			methodBox = color.New(color.BgRed, color.FgWhite).Sprintf(" DELETE ")
		default:
			methodBox = color.New(color.BgWhite, color.FgBlack).Sprintf(" %s ", method)
		}

		// Status code coloring
		switch {
		case status >= 200 && status < 300:
			statusBox = color.New(color.BgGreen, color.FgWhite).Sprintf(" %d ", status)
		case status >= 300 && status < 400:
			statusBox = color.New(color.BgBlue, color.FgWhite).Sprintf(" %d ", status)
		case status >= 400 && status < 500:
			statusBox = color.New(color.BgYellow, color.FgBlack).Sprintf(" %d ", status)
		case status >= 500:
			statusBox = color.New(color.BgRed, color.FgWhite).Sprintf(" %d ", status)
		default:
			statusBox = color.New(color.BgWhite, color.FgBlack).Sprintf(" %d ", status)
		}

		// Format the log message with boxes
		logMessage := fmt.Sprintf("%s %s %s %s - %v",
			color.New(color.FgHiWhite).Sprintf("%-15s", clientIP),
			methodBox,
			statusBox,
			color.New(color.FgHiCyan).Sprintf("%-30s", path),
			color.New(color.FgHiMagenta).Sprintf("%v", latency),
		)

		// Print the log
		fmt.Println(logMessage)
	}
}

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Allow specific origins or use "*" for all origins (less secure)
		// For development, you might want to allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// For production, specify your exact origins:
		// origin := c.Request.Header.Get("Origin")
		// allowedOrigins := []string{
		//     "http://localhost:3000",
		//     "http://localhost:8080",
		//     "https://yourdomain.com",
		// }
		// if isAllowedOrigin(origin, allowedOrigins) {
		//     c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// }

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}

// Helper function to check if origin is allowed (for production use)
func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

// Alternative: More configurable CORS middleware
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// CORSMiddlewareWithConfig creates a CORS middleware with custom configuration
func CORSMiddlewareWithConfig(config CORSConfig) gin.HandlerFunc {
	// Set defaults if not provided
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = []string{"*"}
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	}
	if len(config.AllowHeaders) == 0 {
		config.AllowHeaders = []string{"Content-Type", "Authorization", "X-Requested-With"}
	}
	if config.MaxAge == 0 {
		config.MaxAge = 86400 // 24 hours
	}

	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Set allowed origins
		if len(config.AllowOrigins) == 1 && config.AllowOrigins[0] == "*" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if isAllowedOrigin(origin, config.AllowOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Set other CORS headers
		if config.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", joinStrings(config.AllowMethods))
		c.Writer.Header().Set("Access-Control-Allow-Headers", joinStrings(config.AllowHeaders))
		c.Writer.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", config.MaxAge))

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}

// Helper function to join string slice
func joinStrings(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += ", " + strs[i]
	}
	return result
}
