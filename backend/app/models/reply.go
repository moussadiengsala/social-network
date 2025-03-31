package models

import "time"

type CreateReply struct {
	ID       int    `json:"id" db:"id"`
	Files    string `json:"file" db:"image"`
	Content  string `json:"content" db:"content" validate:"required,min=4,max=255"`
	AuthorID int    `json:"author_id" db:"author_id" validate:"required"`
	EntrieID string `json:"entrie_id" db:"entrie_id" validate:"required"`
}

type GetReply struct {
	ID       int    `json:"id"`
	Files    string `json:"file"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
	EntrieID int    `json:"entrie_id"`

	Likes    []string `json:"likes"`
	DisLikes []string `json:"dis_likes"`

	Reply []string `json:"reply"`

	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`

	CreationDate time.Time
}
