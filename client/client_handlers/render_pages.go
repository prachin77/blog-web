package client_handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prachin77/blog-web/utils"
)

func DefaultRoute(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	// get session token from cookie
	sessionToken, err := utils.GetSessionTokenFromCookie(ctx.Request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error getting session token: %v", err)
		return
	}
	if sessionToken == "" {
		fmt.Println("No session token found, rendering auth page")
		RenderAuthPage(ctx)
		return
	}
	RenderHomePage(ctx, sessionToken)
}

func RenderAuthPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/auth.html")
}

func RenderHomePage(ctx *gin.Context, sessionToken string) {
	ctx.Header("Content-Type", "text/html")
	http.ServeFile(ctx.Writer, ctx.Request, "P:/blog-web/client/templates/home.html")
}
