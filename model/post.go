package model

import "time"

// Post represents a blog article
type Post struct {
	ID       int
	Title    string
	Author   *User
	CreateAt time.Time
	Category *Category
	Content  []byte
	Comments []*Comment
}

// Comment represents a comment of a post
type Comment struct {
	ID       int
	Author   string
	CreateAt time.Time
	Content  []byte
	PostID   int
}

// Category represents blog post category
type Category string
