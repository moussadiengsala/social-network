package lib

import (
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func IsNotificationExist(r *http.Request, notifiction models.Notifications) bool {
	var exist bool
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	query := `
		SELECT EXISTS (
			SELECT 1 FROM notifications WHERE sender_id = ? AND receiver_id = ? AND notification_type = ? AND group_id = ? AND event_id = ?
		)`

	if err := sqlService.SelectSingle(query, []interface{}{notifiction.SenderID, notifiction.ReceiverID, notifiction.NotificationType, notifiction.GroupID, notifiction.EventsID}, &exist); err != nil {
		return exist
	}

	return exist
}
