package models

import "time"

type Session struct {
	Token          string
	ExpirationDate time.Time
	UserID         string
	CreationDate   time.Time
}
