package client_handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prachin77/blog-web/pb"
	"github.com/prachin77/blog-web/utils"
)

var blog_service_client pb.BlogServiceClient


func InitializeBlogClient(client pb.BlogServiceClient) {
	blog_service_client = client
	if client != nil {
		println("✅ BlogServiceClient initialized successfully")
	} else {
		println("❌ BlogServiceClient initialization failed - client is nil")
	}
}

func IsBlogClientInitialized() bool {
	return blog_service_client != nil
}

func CreateBlog(c *gin.Context) {
	if blog_service_client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Blog service client not initialized",
		})
		return
	}

	// Parse form data
	title := c.PostForm("title")
	content := c.PostForm("blog_content")
	authorID := c.PostForm("author_id")
	tag := c.PostForm("tag")
    tag, err := utils.ValidateAndNormalizeTag(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if title == "" || content == "" || authorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, content, and author ID are required",
		})
		return
	}

	// Handle file upload 
	var base64Image string
	file, header, err := c.Request.FormFile("blog_image")
	if err == nil {
		defer file.Close()

		// Check file size (limit to 5MB for Firestore)
		if header.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Image file too large (max 5MB)",
			})
			return
		}

		// Validate image type using utility function
		contentType := header.Header.Get("Content-Type")
		if !utils.IsValidImageType(contentType) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid image type. Only JPEG, JPG, PNG, WEBP, and GIF are allowed",
			})
			return
		}

		// Read file content
		imageBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read image file",
			})
			return
		}

		base64Image = utils.ConvertImageToBase64(imageBytes)
	}

	// set comments and likes to 0 here only , not using postman 
	likes, err := strconv.Atoi(c.PostForm("likes"))
	if err != nil {
		likes = 0 
	}

	comments, err := strconv.Atoi(c.PostForm("comments"))
	if err != nil {
		comments = 0 
	}

	request := &pb.CreateBlogRequest{
		Title:       title,
		BlogContent: content,
		AuthorId:    authorID,
		CreatedAt:   time.Now().Format(time.RFC3339),
		BlogImage:   base64Image, // Base64 string instead of bytes
		Tag:         tag,
		Likes:       int32(likes),
		Comments:    int32(comments),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := blog_service_client.CreateBlog(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to create blog: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     response.Message,
		"blog_id":     response.BlogId,
		"author_id":   response.AuthorId,
		"author_name": response.AuthorName,
	})
}