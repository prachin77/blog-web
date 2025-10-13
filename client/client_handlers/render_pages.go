package client_handlers

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func DefaultRoute(ctx *gin.Context) {
    currentDir, err := os.Getwd()
    if err != nil {
        log.Println("Error getting current directory:", err)
        ctx.String(http.StatusInternalServerError, "Error getting current directory")
        return
    }

    // Navigate to the root directory (going up two levels)
	rootDir := filepath.Dir(currentDir)

    indexFilePath := filepath.Join(rootDir, "client", "templates", "index.html")

    ctx.Header("Content-Type", "text/html")
    http.ServeFile(ctx.Writer, ctx.Request, indexFilePath)
}

func RenderInitPage(ctx *gin.Context) {
    ctx.Header("Content-Type", "text/html")
    RenderAuthPage(ctx)
}

func RenderAuthPage(ctx *gin.Context) {
    currentDir, err := os.Getwd()
    if err != nil {
        log.Println("Error getting current directory:", err)
        ctx.String(http.StatusInternalServerError, "Error getting current directory")
        return
    }

    // Navigate to the root directory (going up two levels))
	rootDir := filepath.Dir(currentDir)

    // Build the relative path to the 'auth.html' file
    authFilePath := filepath.Join(rootDir, "client", "templates", "auth.html")

    tmpl := template.Must(template.ParseFiles(authFilePath))
    if tmpl == nil {
        log.Println("Error: Template loading failed")
        ctx.String(http.StatusInternalServerError, "Error loading template")
        return
    }

    err = tmpl.Execute(ctx.Writer, "message")
    if err != nil {
        log.Println("Error in tmpl.Execute() in RenderAuthPage:", err)
        fmt.Fprint(ctx.Writer, err)
    }
}
