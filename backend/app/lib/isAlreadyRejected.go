package lib

import (
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func IsAlreadyRejectedt(r *http.Request, followerID, followingID int) bool {
	var exist bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT 1 FROM follower WHERE follower_id = ? AND following_id = ? AND status = 'reject'
		)`

	if err := sqlService.SelectSingle(query, []interface{}{followerID, followingID}, &exist); err != nil {
		return exist
	}

	return exist
}
