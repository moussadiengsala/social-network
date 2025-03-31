package models

type Follow struct {
	ID          int    `json:"id" db:"id"`
	FollowerID  int    `json:"follower_id" db:"follower_id" validate:"required"`
	FollowingID int    `json:"following_id" db:"following_id" validate:"required"`
	Status      string `json:"status" db:"status" validate:"required,value=pending|accept|reject"`
}

type GetFollow struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Avatar     string `json:"avatar"`
	Created_at string `json:"created_at"`
}
