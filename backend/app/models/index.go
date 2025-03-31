package models

import (
	"time"
)

type ContextKey string

const (
	UserIDKey     ContextKey = "userID"
	DBInstanceKey ContextKey = "db"
	// Add more custom keys as needed
)

type Comm struct {
	ID             int       `json:"comment_id"`
	Content        string    `json:"content"`
	AuthorID       int       `json:"author_id"`
	AuthorUsername string    `json:"author_username"`
	AuthorAvatar   string    `json:"author_avatar"`
	CreatedAt      time.Time `json:"created_at"`
	LikeCount      int       `json:"like_count"`
	DislikeCount   int       `json:"dislike_count"`
	CommentCount   int       `json:"comment_count"`
	LikeStatus     bool      `json:"like_status"`
	DisLikeStatus  bool      `json:"dislike_status"`
}

type Comment struct {
	ID       string
	PostID   string `json:"post_id"`
	AuthorID string `json:"user_id"`
	Content  string `json:"content"`
}

type RetreivedComment struct {
	ID             string    `json:"comment_id"`
	PostID         string    `json:"post_id"`
	AuthorID       string    `json:"user_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	AuthorAvatar   string    `json:"author_avatar"`
	AuthorUsername string    `json:"author_username"`
	LikeCount      int       `json:"like_count"`
	DislikeCount   int       `json:"dislike_count"`
	LikeStatus     bool      `json:"like_status"`
	DisLikeStatus  bool      `json:"dislike_status"`
}
