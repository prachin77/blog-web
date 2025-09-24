package client_handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prachin77/blog-web/pb"
)

func GetAllBlogs(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	resp, err := blog_service_client.GetAllBlogs(context.Background(), &pb.GetAllBlogsRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	ctx.JSON(http.StatusOK, resp.Blogs)
}
