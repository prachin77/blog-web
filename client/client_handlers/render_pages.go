package client_handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	RenderAuthPage(ctx)
}

func RenderAuthPage(ctx *gin.Context) {
    ctx.Header("Content-Type", "text/html")
    tmpl := template.Must(template.ParseFiles("P:/blog-web/client/templates/auth.html"))
    if tmpl == nil {
        log.Println("Error: Template loading failed")
    }

    err := tmpl.Execute(ctx.Writer, "message")
    if err != nil {
        log.Println("Error in tmpl.Execute() in RenderAuthPage:", err)
        fmt.Fprint(ctx.Writer, err)
    }
}
