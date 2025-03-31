package models

type User struct {
	ID        int    `json:"id" db:"id"`
	Firstname string `json:"first_name" validate:"required,max=16" db:"first_name"`
	Lastname  string `json:"last_name" validate:"required,max=16" db:"last_name"`
	Email     string `json:"email" validate:"required,email" db:"email"`
	Username  string `json:"username" validate:"required,username,min=4,max=16" db:"username"`
	DateBirth string `json:"date_of_birth" validate:"required,date" db:"date_of_birth"`
	Gender    string `json:"gender" validate:"required,value=male|female" db:"gender"`
	Bio       string `json:"bio" validate:"max=255" db:"bio"`
	Avatar    string `json:"avatar" db:"avatar"`
	Password  string `json:"password" validate:"required,password,min=8" db:"password"`
	Privacy   string `json:"privacy" db:"privacy" validate:"required,value=public|private"`
}

type GetUser struct {
	ID              int         `json:"id"`
	Firstname       string      `json:"first_name"`
	Lastname        string      `json:"last_name"`
	Email           string      `json:"email"`
	Username        string      `json:"username"`
	DateBirth       string      `json:"date_of_birth"`
	Gender          string      `json:"gender"`
	Bio             string      `json:"bio"`
	Avatar          string      `json:"avatar"`
	Password        string      `json:"-"`
	Privacy         string      `json:"privacy"`
	CreateAt        string      `json:"create_at"`
	FollowingStatus interface{} `json:"following_status"`

	Followers  int `json:"followers"`
	Following  int `json:"following"`
	LikesCount int `json:"likes_count"`
	PostCount  int `json:"post_count"`

	Online           int `json:"online"`
	UnReadedMessages int `json:"unread_message"`
	LastActivities   any `json:"last_activity"`
}

type Credentials struct {
	Identifiers string `json:"identifiers" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type SettingsUserProfile struct {
	Avatar  string `json:"avatar"`
	Privacy string `json:"privacy"`
	Bio     string `json:"bio"`
}
