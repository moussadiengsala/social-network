package service

import (
	"log"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

func LoggerMiddleware (next lib.Handler) lib.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log something before handling the request
		log.Println("Handling request:", r.URL.Path)
		next(w, r)
	}
}