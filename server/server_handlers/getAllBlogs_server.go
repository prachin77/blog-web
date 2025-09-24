package server_handlers

import (
	"context"
	"log"

	"github.com/prachin77/blog-web/db"
	"github.com/prachin77/blog-web/pb"
)

func (s *BlogServer) GetAllBlogs(ctx context.Context, req *pb.GetAllBlogsRequest) (*pb.GetAllBlogsResponse, error) {
	blogs, err := db.GetAllBlogs()
	if err != nil {
		log.Printf("❌ Failed to fetch blogs: %v", err)
		return nil, err
	}

	var blogResponses []*pb.SingleBlog
	for _, blog := range blogs {
		user, err := db.GetUserByID(blog.AuthorID)
		authorName := "Unknown"
		if err == nil {
			authorName = user.Username
		}

		blogResponses = append(blogResponses, &pb.SingleBlog{
			BlogId:     blog.BlogID,
			AuthorId:   blog.AuthorID,
			AuthorName: authorName,
			Title:      blog.Title,
			BlogImage:  blog.BlogImage,
			Tag:        blog.Tags,
			CreatedAt:  blog.CreatedAt.String(),
			Likes:      int32(blog.Likes),
			Comments:   int32(blog.Comments),
		})
	}

	return &pb.GetAllBlogsResponse{
		Blogs: blogResponses,
	}, nil
}
