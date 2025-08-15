// package client_handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func DefaultRoute(ctx *gin.Context) {
// 	ctx.Header("Content-Type", "text/html")
// 	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/index.html")
// }

// func RenderInitPage(ctx *gin.Context) {
// 	// this func checks whether there are cookies or not
// 	// if no cookie found then again auth page

// 	ctx.Header("Content-Type", "text/html")
// 	RenderHomePage(ctx)
// }

// // RenderAuthPage serves the auth fragment directly
// func RenderAuthPage(ctx *gin.Context) {
// 	ctx.Header("Content-Type", "text/html")
// 	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/auth.html")
// }

// // RenderHomePage serves the logged-in home fragment
// func RenderHomePage(ctx *gin.Context) {
// 	ctx.Header("Content-Type", "text/html")
// 	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/home.html")
// }

package client_handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	// Use absolute path
	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	// this func checks whether there are cookies or not
	// if no cookie found then again auth page

	ctx.Header("Content-Type", "text/html")
	// Add HTMX-specific headers
	ctx.Header("HX-Trigger", "page-loaded")
	RenderHomePage(ctx)
}

// RenderAuthPage serves the auth fragment directly
func RenderAuthPage(ctx *gin.Context) {
	// Set proper headers for HTMX
	ctx.Header("Content-Type", "text/html")
	ctx.Header("HX-Trigger", "auth-page-loaded")

	// Read and serve the auth.html file content
	authPath := "P:/blog-web/client/templates/auth.html"
	

	// Check if file exists
	if _, err := os.Stat(authPath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Auth template not found"})
		return
	}

	// Read file content
	file, err := os.Open(authPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open auth template"})
		return
	}
	defer file.Close()

	// Copy file content to response
	ctx.Status(http.StatusOK)
	io.Copy(ctx.Writer, file)
}

// RenderHomePage serves the logged-in home fragment
func RenderHomePage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.Header("HX-Trigger", "home-page-loaded")

	// Read and serve the home.html file content
	homePath := "P:/blog-web/client/templates/home.html"

	// Check if file exists
	if _, err := os.Stat(homePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Home template not found"})
		return
	}

	// Read file content
	file, err := os.Open(homePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open home template"})
		return
	}
	defer file.Close()

	// Copy file content to response
	ctx.Status(http.StatusOK)
	io.Copy(ctx.Writer, file)
}

// Alternative: Using embedded content (if you want to embed HTML directly)
func RenderAuthPageDirect(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.Header("HX-Trigger", "auth-page-loaded")

	authHTML := `<div id="main_auth_div" class="min-h-screen flex items-center justify-center bg-gradient-to-r from-indigo-100 via-purple-100 to-pink-100 p-6">
		<div class="bg-white rounded-2xl shadow-xl w-full max-w-4xl grid md:grid-cols-2 overflow-hidden">
			<!-- Your auth form content here -->
		</div>
	</div>`

	ctx.Data(http.StatusOK, "text/html", []byte(authHTML))
}
