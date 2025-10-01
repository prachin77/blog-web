package server_handlers

import (
	"context"
	"fmt"

	"github.com/prachin77/blog-web/db"
	"github.com/prachin77/blog-web/pb"
)

func (s *BlogServer) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.BlogResponse, error) {
	blogID := req.GetBlogId()
	authorID := req.GetAuthorId()

	if blogID == "" || authorID == "" {
		return nil, fmt.Errorf("blog_id and author_id are required")
	}

	// Verify blog exists and belongs to the author
	blog, err := db.GetBlogByID(blogID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch blog: %w", err)
	}

	if blog.AuthorID != authorID {
		return nil, fmt.Errorf("unauthorized: blog does not belong to author_id %s", authorID)
	}

	if err := db.DeleteBlog(blogID); err != nil {
		return nil, fmt.Errorf("failed to delete blog: %w", err)
	}

	if err := db.DecrementUserBlogCount(authorID); err != nil {
		fmt.Printf("❌ Failed to decrement blog count for user %s", authorID)
	}

	fmt.Printf("✅ Deleted blog with ID %s by author %s", blogID, authorID)

	return &pb.BlogResponse{
		Message:  "Blog deleted successfully",
		BlogId:   blogID,
		AuthorId: blog.AuthorID,
	}, nil
}
