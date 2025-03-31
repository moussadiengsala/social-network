package lib

import (
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func GetAllGroupsMembers(r *http.Request, groupID, userID int, isRequestedUserNeeded bool) ([]int, error) {
	var members []int
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	extraCondition := ""
	if !isRequestedUserNeeded {
		extraCondition = fmt.Sprintf("AND user_id != %d", userID)
	}

	query := fmt.Sprintf(`SELECT user_id FROM groupMembers WHERE group_id = "%d" AND status = "accepted" %s`, groupID, extraCondition)
	err := sqlService.SelectAllForPrimitive(query, []interface{}{}, &members)
	if err != nil {
		return nil, fmt.Errorf("failed to get user's friends: %v", err)
	}

	return members, nil
}
