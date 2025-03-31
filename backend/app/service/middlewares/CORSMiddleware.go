package service

import (
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

// Implements a CORS middleware

func CORSMiddleware(next lib.Handler) lib.Handler {
	fmt.Println("cors controller middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		next(w, r)
	}
}
