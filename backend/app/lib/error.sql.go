package lib

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/mattn/go-sqlite3"
)

func SqlError(err error, lookingForFields []string, foreignFields []string) (string, int) {
	if sqliteErr, ok := err.(sqlite3.Error); ok {

		switch sqliteErr.ExtendedCode {
		case sqlite3.ErrConstraintForeignKey:
			return fmt.Sprintf("The %s must exist to perform this operation.", strings.Join(foreignFields, " or ")), http.StatusBadRequest
		case sqlite3.ErrConstraintPrimaryKey, sqlite3.ErrConstraintUnique:
			return fmt.Sprintf("The %s already exist.", strings.Join(lookingForFields, " or ")), http.StatusConflict
		case sqlite3.ErrConstraintCheck:
			return "Trouble to perform this operations with the values youprovides.", http.StatusBadRequest
		default:
			return fmt.Sprintf("SQLite error: %s", err.Error()), http.StatusInternalServerError
		}
	}

	if err == sql.ErrNoRows {
		return fmt.Sprintf("%s not found.", strings.Join(lookingForFields, " or ")), http.StatusNotFound
	}

	return "Something went wrong. If the problem persists, please contact us.", http.StatusInternalServerError
}
