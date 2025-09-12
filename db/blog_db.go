package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/prachin77/blog-web/model"
)


// CreateBlog saves a new blog to Firestore
func CreateBlog(blog *model.Blog) (*model.Blog, error) {
	ctx := context.Background()

	// Generate a unique blog ID
	blogRef := FirestoreClient.Collection(BlogsCollection).NewDoc()
	blog.BlogID = blogRef.ID
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	// Prepare data for Firestore
	blogData := map[string]interface{}{
		"blog_id":      blog.BlogID,
		"title":        blog.Title,
		"blog_content": blog.BlogContent,
		"author_id":    blog.AuthorID,
		"created_at":   blog.CreatedAt,
		"updated_at":   blog.UpdatedAt,
		"tags":         blog.Tags,
		"blog_image":   blog.BlogImage, // Base64 string
		"likes":        blog.Likes,
		"comments":     blog.Comments,
	}

	// Save blog to Firestore
	_, err := blogRef.Set(ctx, blogData)
	if err != nil {
		return nil, fmt.Errorf("failed to create blog: %w", err)
	}

	log.Printf("✅ Blog created with ID: %s", blog.BlogID)
	return blog, nil
}


// IncrementUserBlogCount increases the blog count for a user
func IncrementUserBlogCount(userID string) error {
	ctx := context.Background()

	userRef := FirestoreClient.Collection("users").Doc(userID)
	_, err := userRef.Update(ctx, []firestore.Update{
		{
			Path:  "NoOfBlogs",
			Value: firestore.Increment(1),
		},
	})

	if err != nil {
		return fmt.Errorf("failed to increment blog count: %w", err)
	}

	return nil
}