package models

type Groups struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" validate:"required,min=4,max=100"`
	Description string `json:"description" db:"description" validate:"required,min=4,max=255"`
	AuthorID    int    `json:"-" db:"author_id" validate:"required,numeric"`
	Cover       string `json:"cover" db:"cover"`
}

type GroupMembers struct {
	Status  string `json:"status" db:"status" validate:"required,value=invited|accepted|requested"`
	GroupID int    `json:"group_id" db:"group_id" validate:"required,numeric"`
	User    string `json:"user" db:"user_id" validate:"required,selected_user"`
	Role    string `json:"-" db:"role" valient:"required,value=admin|user"`
}

type GetGroupMembers struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
}

type GroupPost struct {
	ID       int    `json:"id" db:"id"`
	Image    string `json:"file" db:"image"`
	Content  string `json:"content" validate:"required,min=4,max=255" db:"content"`
	GroupID  int    `json:"group_id" validate:"required,numeric" db:"group_id"`
	AuthorID int    `json:"author_id" validate:"required,numeric" db:"author_id"`
}

type GroupComments struct {
	ID       int    `json:"id" db:"id"`
	Image    string `json:"file" db:"image"`
	Content  string `json:"content" validate:"required,min=4,max=255" db:"content"`
	GroupID  int    `json:"group_id" validate:"required,numeric" db:"group_id"`
	AuthorID int    `json:"author_id" validate:"required,numeric" db:"author_id"`
	PostID   string `json:"post_id" validate:"required,numeric" db:"post_id"`
}

type GroupEvents struct {
	ID          int    `json:"id" db:"id"`
	GroupID     int    `json:"group_id" db:"group_id" validate:"required,numeric"`
	Title       string `json:"title" db:"title" validate:"required,min=4,max=100"`
	Description string `json:"description" db:"description" validate:"required,min=4,max=255"`
	AuthorID    int    `json:"author_id" validate:"required,numeric" db:"author_id"`
	DateTime    string `json:"datetime" validate:"required,date" db:"datetime"`
}

type GetGroupEvents struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    int    `json:"author_id"`
	DateTime    string `json:"datetime"`

	AuthorFirstname string `json:"author_first_name"`
	AuthorLastname  string `json:"author_last_name"`
	AuthorEmail     string `json:"author_email"`
	AuthorUsername  string `json:"author_username"`
	AuthorAvatar    string `json:"author_avatar"`

	ResponseStatus string `json:"response_status"`
	GoingCount     int    `json:"going_count"`
	NotGoingCount  int    `json:"not_going_count"`
}

type GroupEventResponses struct {
	Response string `json:"response" db:"response" validate:"required,value=going|not going"`
	EventID  int    `json:"event_id" db:"event_id" validate:"required,numeric"`
	GroupID  int    `json:"group_id" db:"group_id" validate:"required,numeric"`
	UserID   int    `json:"user_id" validate:"required,numeric" db:"user_id"`
}

type GetGroupEventResponses struct {
	Response  string `json:"response"`
	UserID    int    `json:"user_id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
}

type GetGroups struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    int    `json:"author_id"`
	Cover       string `json:"cover"`
	CreateAt    string `json:"created_at"`

	AuthorUsername  string `json:"author_username"`
	AuthorFirstname string `json:"author_first_name"`
	AuthorLastname  string `json:"author_last_name"`
	AutherAvatar    string `json:"author_avatar"`

	MemberCount   int    `json:"member_count"`
	PostCount     int    `json:"post_count"`
	EventCount    int    `json:"event_count"`
	UnReadMessage int    `json:"unread_message"`
	LastActivity  string `json:"last_activity"`

	IsPartOfGroup int    `json:"is_part_of_group"`
	Status        string `json:"status"`
}
