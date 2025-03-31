package models

import "time"

type Notifications struct {
	ID               int    `json:"id" db:"id"`
	SenderID         int    `json:"sender_id" db:"sender_id" validate:"required"`
	ReceiverID       int    `json:"receiver_id" db:"receiver_id" validate:"required"`
	GroupID          int    `json:"group_id" db:"group_id"`
	EventsID         int    `json:"events_id" db:"event_id"`
	NotificationType string `json:"notification_type" db:"notification_type" validate:"required,value=follow_request|message|groups_invited|groups_requested|groups_events"`
}

type GetNotifications struct {
	ID               int       `json:"id" db:"id"`
	SenderID         int       `json:"sender_id" db:"sender_id"`
	ReceiverID       int       `json:"receiver_id" db:"receiver_id"`
	SenderUsername   string    `json:"sender_username" db:"sender_username"`
	NotificationType string    `json:"notification_type" db:"notification_type"`
	GroupID          int       `json:"group_id" db:"group_id"`
	Status           string    `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type GetPrivateNotification struct {
	ID               int    `json:"id"`
	SenderID         int    `json:"sender_id"`
	ReceiverID       int    `json:"receiver_id"`
	NotificationType string `json:"notification_type"`

	SenderFirstname string `json:"sender_first_name"`
	SenderLastname  string `json:"sender_last_name"`
	SenderUsernase  string `json:"sender_username"`
	SenderAvatar    string `json:"sender_avatar"`

	ReceiverFirstname string `json:"receiver_first_name"`
	ReceiverLastname  string `json:"receiver_last_name"`
	ReceiverUsername  string `json:"receiver_username"`
	ReceiverAvatar    string `json:"receiver_avatar"`
}

type GetGroupsInvitationRequestNotification struct {
	ID               int    `json:"id"`
	SenderID         int    `json:"sender_id"`
	ReceiverID       int    `json:"receiver_id"`
	NotificationType string `json:"notification_type"`

	SenderFirstname string `json:"sender_first_name"`
	SenderLastname  string `json:"sender_last_name"`
	SenderUsernase  string `json:"sender_username"`
	SenderAvatar    string `json:"sender_avatar"`

	ReceiverFirstname string `json:"receiver_first_name"`
	ReceiverLastname  string `json:"receiver_last_name"`
	ReceiverUsername  string `json:"receiver_username"`
	ReceiverAvatar    string `json:"receiver_avatar"`

	GroupID     int    `json:"group_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetGroupsEventsNotification struct {
	ID               int    `json:"id"`
	SenderID         int    `json:"sender_id"`
	ReceiverID       int    `json:"receiver_id"`
	NotificationType string `json:"notification_type"`

	SenderFirstname string `json:"sender_first_name"`
	SenderLastname  string `json:"sender_last_name"`
	SenderUsernase  string `json:"sender_username"`
	SenderAvatar    string `json:"sender_avatar"`

	ReceiverFirstname string `json:"receiver_first_name"`
	ReceiverLastname  string `json:"receiver_last_name"`
	ReceiverUsername  string `json:"receiver_username"`
	ReceiverAvatar    string `json:"receiver_avatar"`

	GroupID           int    `json:"group_id"`
	GroupsTitle       string `json:"group_title"`
	GroupsDescription string `json:"group_description"`

	EventID           int    `json:"event_id"`
	EventsTitle       string `json:"event_title"`
	EventsDescription string `json:"event_description"`
	AuthorID          int    `json:"author_id"`
	DateTime          string `json:"datetime"`

	AuthorFirstname string `json:"author_first_name"`
	AuthorLastname  string `json:"author_last_name"`
	AuthorUsernase  string `json:"author_username"`
	AuthorAvatar    string `json:"author_avatar"`
}
