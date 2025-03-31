package lib

import (
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func GetUsersFriends(r *http.Request, userID int) ([]int, error) {
	var friendIDs []int
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
        SELECT follower_id FROM follower WHERE (following_id = ?) AND status = 'accept'
		UNION
		SELECT following_id FROM follower WHERE (follower_id = ?) AND status = 'accept'
    `

	err := sqlService.SelectAllForPrimitive(query, []interface{}{userID, userID}, &friendIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get user's friends: %v", err)
	}

	return friendIDs, nil
}
