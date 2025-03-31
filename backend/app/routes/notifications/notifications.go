package notifications

import (
	"fmt"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

type Notifications struct {
	App *core.App
}

func (n *Notifications) Route(app *core.App) {
	n.App = app
	app.GET("/notifications", n.GetAllNotification)
}

func (n Notifications) CreateNotification(r *http.Request, response *lib.Response, credentials *models.Notifications) {

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	notification := lib.CreateData{
		ID:            0,
		Credentials:   *credentials,
		Table:         "notifications",
		ForeignFields: []string{"user"},
	}

	notification.Create(response, sqlService)
	credentials.ID = notification.ID
}

func (n Notifications) MarkAsReaded(r *http.Request, response *lib.Response, senderID, receiverID int, notificationType string) {
	if response.Code != 200 {
		return
	}
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(`UPDATE notifications SET status = "read" WHERE sender_id = %d AND receiver_id = %d AND notification_type = '%s';`, senderID, receiverID, notificationType)
	if _, insertErr := sqlService.Update(query); insertErr != nil {
		response.Data = nil
		message, statusCode := lib.SqlError(insertErr, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(response, message, statusCode)
		return
	}
}

func (n Notifications) CleanNotification(r *http.Request, response *lib.Response, data ...interface{}) {
	if response.Code != 200 {
		return
	}
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(`DELETE FROM notifications WHERE sender_id = %v AND receiver_id = %v AND notification_type = '%v' AND group_id = %v AND event_id = %v;`, data...)
	if insertErr := sqlService.Delete(query); insertErr != nil {
		response.Data = nil
		message, statusCode := lib.SqlError(insertErr, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(response, message, statusCode)
		return
	}
}

func (n *Notifications) GetAllNotification(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	notificationType := params.Get("notification")
	limit := params.Get("limit")
	offset := params.Get("offset")

	switch notificationType {
	case "groups_events":
		credentials := []models.GetGroupsEventsNotification{}
		notification := lib.GetFeed{
			Credentials:     &credentials,
			CredentialsType: models.GetGroupsEventsNotification{},
			Query:           fmt.Sprintf(internals.QUERY_GETTING_ALL_GROUP_EVENTS_NOTIFICATION, userID, limit, offset),
		}
		notification.GetAllFeed(r, &response)
		// Mark all retrieved notifications as read
		for _, notification := range credentials {
			if notification.ReceiverID == userID {
				n.MarkAsReaded(r, &response, notification.SenderID, notification.ReceiverID, notification.NotificationType)
			}
		}
	case "groups_invited", "groups_requested":
		credentials := []models.GetGroupsInvitationRequestNotification{}
		notification := lib.GetFeed{
			Credentials:     &credentials,
			CredentialsType: models.GetGroupsInvitationRequestNotification{},
			Query:           fmt.Sprintf(internals.QUERY_GETTING_ALL_GROUPS_INVITED_REQUESTED_NOTIFICATION, userID, limit, offset),
		}
		notification.GetAllFeed(r, &response)
		// Mark all retrieved notifications as read
		for _, notification := range credentials {
			if notification.ReceiverID == userID {
				n.MarkAsReaded(r, &response, notification.SenderID, notification.ReceiverID, notification.NotificationType)
			}
		}
	case "follow_request":
		credentials := []models.GetPrivateNotification{}
		notification := lib.GetFeed{
			Credentials:     &credentials,
			CredentialsType: models.GetPrivateNotification{},
			Query:           fmt.Sprintf(internals.QUERY_GETTING_ALL_FOLLOW_REQUEST_NOTIFICATION, userID, limit, offset),
		}
		notification.GetAllFeed(r, &response)
		// Mark all retrieved notifications as read
		for _, notification := range credentials {
			if notification.ReceiverID == userID {
				n.MarkAsReaded(r, &response, notification.SenderID, notification.ReceiverID, notification.NotificationType)
			}
		}
	default:
		lib.ErrorWriter(&response, "unknown notification", http.StatusBadRequest)
	}

	client, ok := n.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadUnreadNotification(r, *n.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATEUNREADNOTIFICATION",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userID})
	}

	lib.ResponseFormatter(w, response)
}
