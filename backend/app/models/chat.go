package models

import "time"

type PrivateMessage struct {
	ID         int    `json:"id" db:"id"`
	SenderID   int    `json:"sender_id" db:"sender_id" validate:"required,numeric"`
	ReceiverID int    `json:"receiver_id" db:"receiver_id" validate:"required,numeric"`
	Content    string `json:"content" db:"content" validate:"required"`
}

type GroupMessage struct {
	ID       int    `json:"id" db:"id"`
	SenderID int    `json:"sender_id" db:"sender_id" validate:"required,numeric"`
	GroupID  int    `json:"group_id" db:"group_id" validate:"required,numeric"`
	Content  string `json:"content" db:"content" validate:"required"`
}

type GetPrivateMessage struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`

	SenderFirstname string `json:"sender_first_name"`
	SenderLastname  string `json:"sender_last_name"`
	SenderUsernase  string `json:"sender_username"`
	SenderAvatar    string `json:"sender_avatar"`

	ReceiverFirstname string `json:"receiver_first_name"`
	ReceiverLastname  string `json:"receiver_last_name"`
	ReceiverUsername  string `json:"receiver_username"`
	ReceiverAvatar    string `json:"receiver_avatar"`

	CreatedAt time.Time `json:"created_at"`
}

type GetGroupMessage struct {
	ID        int       `json:"id"`
	SenderID  int       `json:"sender_id"`
	GroupID   int       `json:"group_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	SenderFirstname string `json:"sender_first_name"`
	SenderLastname  string `json:"sender_last_name"`
	SenderUsernase  string `json:"sender_username"`
	SenderAvatar    string `json:"sender_avatar"`

	GroupTitle       string `json:"group_title"`
	GroupDescription string `json:"group_description"`
	GroupCover       string `json:"group_cover"`
	GroupCreateAt    string `json:"group_created_at"`
}
