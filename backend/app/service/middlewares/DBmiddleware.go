package service

import (
	"context"
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	migrate "learn.zone01dakar.sn/forum-rest-api/pkg/db/sqlite"
)

func DBMiddleware(next lib.Handler) lib.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		response := lib.Response{}
		DB := lib.DB{Instance: nil, Err: nil}

		database := migrate.Config{
			Driver: "sqlite3",
			Name:   "forum.db",
		}
		/* Check if a database connection is already opened
		to avoid opening multiple databases connections */
		if DB.Instance == nil {
			DB.Instance, DB.Err = database.Inits()
			// internals.DropAnTable(DB.Instance, "follow")
			// Throw the error page when we have a database issue
			if DB.Err != nil {
				fmt.Println("DB.Err", DB.Err)
				lib.ErrorWriter(&response, DB.Err.Error(), http.StatusInternalServerError)
				lib.ResponseFormatter(w, response)
				return
			}
			//internals.TablesCreation(DB.Instance)
		}

		ctx := context.WithValue(r.Context(), models.DBInstanceKey, DB)
		next(w, r.WithContext(ctx))
	}
}
