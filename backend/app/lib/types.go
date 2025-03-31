package lib

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type Middleware func(Handler) Handler

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Payload struct {
	User           models.GetUser `json:"user"`
	ExpirationDate time.Time      `json:"expiration"`
}

func (p Payload) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

type DB struct {
	Instance *sql.DB
	Err      error
}
