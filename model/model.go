package model

import "time"

type User struct {
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	NoOfBlogs  int       `json:"no_of_blogs"`
	Followers  int       `json:"followers"`
	Followings int       `json:"followings"`
}

// using entire "User" struct in type which
// will expose more data than required making
// operations slow , but to ease things up i'm doin now
// [not reccommended for industry level]

type Followers struct {
	FollowersList []User `json:"followers_list"`
	FollowingList []User `json:"following_list"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Blog struct {
    BlogID      string    `firestore:"blog_id" json:"blog_id"`
    Title       string    `firestore:"title" json:"title"`
    BlogContent string    `firestore:"blog_content" json:"blog_content"`
    AuthorID    string    `firestore:"author_id" json:"author_id"`
    CreatedAt   time.Time `firestore:"created_at" json:"created_at"`
    UpdatedAt   time.Time `firestore:"updated_at" json:"updated_at"`
    Tags        string    `firestore:"tags" json:"tags"`
    BlogImage   string    `firestore:"blog_image" json:"blog_image"`
    Likes       int       `firestore:"likes" json:"likes"`
    Comments    int       `firestore:"comments" json:"comments"`
}