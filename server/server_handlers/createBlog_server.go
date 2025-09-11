package server_handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prachin77/blog-web/db"
	"github.com/prachin77/blog-web/model"
	"github.com/prachin77/blog-web/pb"
)

type BlogServer struct {
	pb.UnimplementedBlogServiceServer
}

func (s *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.BlogResponse, error) {
	log.Printf("📝 Creating blog: %s", req.Title)

	// Validate request
	if req.Title == "" || req.BlogContent == "" || req.AuthorId == "" {
		return nil, fmt.Errorf("title, content, and author ID are required")
	}

	var createdAt time.Time
	if req.CreatedAt != "" {
		parsedTime, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			log.Printf("⚠️ Failed to parse created_at time, using current time: %v", err)
			createdAt = time.Now()
		} else {
			createdAt = parsedTime
		}
	} else {
		createdAt = time.Now()
	}

	author, err := db.GetUserByID(req.AuthorId)
	if err != nil {
		log.Printf("❌ Author not found: %v", err)
		return nil, fmt.Errorf("author not found: %v", err)
	}

	blog := &model.Blog{
		Title:       req.Title,
		BlogContent: req.BlogContent,
		AuthorID:    req.AuthorId,
		CreatedAt:   createdAt,
		UpdatedAt:   time.Now(),
		Tags:        req.Tag,
		BlogImage:   req.BlogImage, // Base64 string from request
		Likes:       int(req.Likes),
		Comments:    int(req.Comments),
	}

	// Save blog to database
	createdBlog, err := db.CreateBlog(blog)
	if err != nil {
		log.Printf("❌ Failed to create blog: %v", err)
		return nil, fmt.Errorf("failed to create blog: %v", err)
	}

	// Update author's blog count
	err = db.IncrementUserBlogCount(req.AuthorId)
	if err != nil {
		log.Printf("⚠️ Failed to update author blog count: %v", err)
		// Don't fail the entire operation for this
	}

	log.Printf("✅ Blog created successfully: %s", createdBlog.BlogID)

	return &pb.BlogResponse{
		Message:    "Blog created successfully",
		BlogId:     createdBlog.BlogID,
		AuthorId:   createdBlog.AuthorID,
		AuthorName: author.Username,
	}, nil
}