package models

import "time"

type CreateComments struct {
	ID       int    `json:"id" db:"id"`
	Image    string `json:"file" db:"image"`
	Content  string `json:"content" db:"content" validate:"required,min=4,max=255"`
	AuthorID int    `json:"author_id" db:"author_id" validate:"required"`
	EntrieID int    `json:"entrie_id" db:"entrie_id" validate:"required"`
}

type GetComments struct {
	ID             int       `json:"id" db:"id"`
	EntrieID       int       `json:"entrie_id" db:"entrie_id"`
	AuthorID       int       `json:"author_id" db:"author_id"`
	Content        string    `json:"content" db:"content"`
	Image          string    `json:"file" db:"image"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	AuthorAvatar   string    `json:"avatar" db:"avatar"`
	AuthorUsername string    `json:"username" db:"username"`
	LikeCount      int       `json:"like_count"`
	LikeStatus     bool      `json:"like_status"`
}

type GetGroupComment struct {
	ID        int       `json:"id" db:"id"`
	GroupID   int       `json:"group_id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	Content   string    `json:"content" db:"content"`
	Image     string    `json:"image" db:"image"`
	PostID    int       `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Avatar    string `json:"avatar" db:"avatar"`

	GroupTitle       string `json:"group_title"`
	GroupDescription string `json:"group_description"`
	GroupCover       string `json:"group_cover"`
	GroupCreateAt    string `json:"group_create_at"`
}
