package lib

import (
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func IsMembersExist(r *http.Request, groupID, userID int) bool {
	var exist bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	// AND status = "accepted"
	query := `
		SELECT EXISTS (
			SELECT role FROM groupMembers WHERE group_id = ? AND user_id = ?
		)`

	if insertPostErr := sqlService.SelectSingle(query, []interface{}{groupID, userID}, &exist); insertPostErr != nil {
		return exist
	}

	return exist
}

func IsMembersExistAndAccepted(r *http.Request, groupID, userID int) bool {
	var exist bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT role FROM groupMembers WHERE group_id = ? AND user_id = ? AND status = "accepted"
		)`

	if insertPostErr := sqlService.SelectSingle(query, []interface{}{groupID, userID}, &exist); insertPostErr != nil {
		return exist
	}

	return exist
}
