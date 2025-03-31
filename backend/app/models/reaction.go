package models

import "time"

type Reaction struct {
	ID        int       `json:"id" db:"id"`
	EntriesID int       `json:"entries_id" db:"entries_id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	Action    string    `json:"action" db:"action"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
