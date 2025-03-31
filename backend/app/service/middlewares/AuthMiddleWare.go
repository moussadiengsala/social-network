package service

import (
	"context"
	"net/http"
	"strings"

	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/jwt"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
)

func AuthMiddleware(next lib.Handler) lib.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/api/auth/") || strings.Contains(r.URL.Path, "/api/static/upload") || strings.Contains(r.URL.Path, "/api/ws") {
			next(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")

		j := jwt.JWT{}
		response := lib.Response{}
		payload, err := j.CheckingToken(r, &response, authorizationHeader)

		if err != nil {
			lib.ResponseFormatter(w, response)
			return
		}
		ctx := context.WithValue(r.Context(), models.UserIDKey, payload.User.ID)
		next(w, r.WithContext(ctx))
	}
}
