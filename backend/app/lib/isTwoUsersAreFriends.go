package lib

import (
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func IsTwoUsersAreFreinds(r *http.Request, user1ID, user2ID int) bool {
	var exists bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT 1 FROM follower WHERE (follower_id = '%d' AND following_id = '%d' AND status = 'accept') OR (follower_id = '%d' AND following_id = '%d' AND status = 'accept')
		)`

	err := sqlService.SelectSingle(fmt.Sprintf(query, user1ID, user2ID, user2ID, user1ID), []interface{}{}, &exists)
	if err != nil {
		return exists
	}

	return exists
}

func IsFollowsUser(r *http.Request, followerID, followingID int) bool {
	var exists bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT 1 FROM follower WHERE (follower_id = '%d' AND following_id = '%d' AND status = 'accept')
	)`

	err := sqlService.SelectSingle(fmt.Sprintf(query, followerID, followingID), []interface{}{}, &exists)
	if err != nil {
		return exists
	}

	return exists
}
