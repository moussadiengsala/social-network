package models

import (
	"time"
)

type CreatePost struct {
	ID            int    `json:"id" db:"id"`
	Image         string `json:"photo" db:"image"`
	Content       string `json:"content" validate:"required,min=4,max=255" db:"content"`
	AuthorID      int    `json:"author_id" validate:"required" db:"author_id"`
	Privacy       string `json:"privacy" db:"privacy" validate:"required,value=private|public|almost_private"`
	SelectedUsers string `json:"selected_users" validate:"selected_user"`
}

type GetPost struct {
	ID         int       `json:"id" db:"id"`
	Content    string    `json:"content" db:"content"`
	AuthorID   int       `json:"author_id" db:"author_id"`
	Image      string    `json:"image" db:"image"`
	Privacy    string    `json:"privacy" db:"privacy"`
	Comments   int       `json:"comments" db:"comments"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Username   string    `json:"username" db:"username"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Likes      int       `json:"likes" db:"likes"`
	LikeStatus bool      `json:"like_status" db:"like_status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type PostVisibility struct {
	PostID string `json:"post_id" db:"post_id"`
	UserID string `json:"user_id" db:"user_id"`
}

type GetGroupPost struct {
	ID        int       `json:"id" db:"id"`
	GroupID   int       `json:"group_id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	Content   string    `json:"content" db:"content"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Avatar    string `json:"avatar" db:"avatar"`

	CommentCount int `json:"comments" db:"comment_count"`

	GroupTitle       string `json:"group_title"`
	GroupDescription string `json:"group_description"`
	GroupCover       string `json:"group_cover"`
	GroupCreateAt    string `json:"group_create_at"`
}
