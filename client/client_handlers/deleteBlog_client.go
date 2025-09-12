package client_handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prachin77/blog-web/pb"
)

type DeleteBlogRequest struct {
	AuthorID string `json:"author_id"`
}

func DeleteBlog(ctx *gin.Context) {
	blogID := ctx.Param("blog_id") 
	if blogID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Blog ID is required"})
		return
	}

	var body DeleteBlogRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Make gRPC request to delete blog
	req := &pb.DeleteBlogRequest{
		BlogId:   blogID,
		AuthorId: body.AuthorID,
	}

	resp, err := blog_service_client.DeleteBlog(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":   resp.Message,
		"blog_id":   resp.BlogId,
		"author_id": resp.AuthorId,
	})
}
