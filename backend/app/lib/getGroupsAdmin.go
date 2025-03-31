package lib

import (
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func GetGroupsAdminMember(r *http.Request, groupID int) int {
	var adminID int
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `SELECT user_id FROM groupMembers WHERE (group_id = ? AND role = 'admin')`

	err := sqlService.SelectSingle(query, []interface{}{groupID}, &adminID)
	if err != nil {
		return 0
	}

	return adminID
}
