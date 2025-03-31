package lib

import (
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func IsPrivateAccount(r *http.Request, response *Response, userID int) bool {
	var exist bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT 1 FROM user WHERE id = ? AND privacy = 'private'
		)`

	if err := sqlService.SelectSingle(query, []interface{}{userID}, &exist); err != nil {
		return exist
	}

	return exist
}
