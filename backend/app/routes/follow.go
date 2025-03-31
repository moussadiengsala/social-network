package routes

import (
	"fmt"
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/notifications"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

type Follow struct {
	App *core.App
}

func (f *Follow) Route(app *core.App) {
	f.App = app

	app.POST("/follow", f.HandleFollowEvent)
	app.GET("/accept-follow", f.AcceptARequestOrReject)

}

func (f Follow) HandleFollowEvent(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	userID := r.Context().Value(models.UserIDKey).(int)

	credentials := models.Follow{}
	credentials.FollowerID = lib.Convert(r.URL.Query().Get("follower_id"))
	credentials.FollowingID = userID
	credentials.Status = "accept"
	isUnfollowed := false

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil || credentials.FollowingID == credentials.FollowerID {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		if credentials.FollowingID == credentials.FollowerID {
			response.Message = fmt.Sprintf("%s * you cannot follow yourself", response.Message)
		}
		lib.ResponseFormatter(w, response)
		return
	}

	// if the user has been already rejected or not.
	if isAlreadyRejected := lib.IsAlreadyRejectedt(r, credentials.FollowerID, credentials.FollowingID); isAlreadyRejected {
		lib.ErrorWriter(&response, "this user reject your request you cannot perform this operation.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	// check if the following_id is private and they are not freind yet to send the notification.
	if isPrivateAccount := lib.IsPrivateAccount(r, &response, credentials.FollowerID); isPrivateAccount {
		credentials.Status = "pending"
	}

	// follow
	f.Follow(r, &response, &credentials)

	// unfollow
	if response.Code == 409 {
		isUnfollowed = true
		f.UnFollow(r, &response, credentials)
	}

	// send notification
	payload, err := utils.GetPayloadUser(r, *f.App.Hub, credentials.FollowerID, &response)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		lib.ResponseFormatter(w, response)
		return
	}

	client, ok := f.App.Hub.Clients[credentials.FollowerID]

	// store the notification
	notification := notifications.Notifications{}
	credentialsNotification := models.Notifications{
		SenderID:         credentials.FollowingID,
		ReceiverID:       credentials.FollowerID,
		NotificationType: "follow_request",
	}

	if isNotificationExist := lib.IsNotificationExist(r, credentialsNotification); !isNotificationExist && !isUnfollowed && credentials.Status == "pending" {
		p := utils.ResponseWS{Users: payload.Users}
		// store the notification in the database
		notification.CreateNotification(r, &response, &credentialsNotification)
		if response.Code != http.StatusOK {
			lib.ResponseFormatter(w, response)
			return
		}

		if ok {
			unreadNotification, err := utils.GetPayloadUnreadNotification(r, *f.App.Hub, credentials.FollowerID, &response)
			if err != nil {
				message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
				lib.ErrorWriter(&response, message, statusCode)
				utils.HandleUnexpectedEvent(client, &response, []int{userID})
				return
			}

			p.UnreadNotifications = unreadNotification.UnreadNotifications
			p.Items = credentialsNotification
			utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
				Action:  utils.NOTIFICATIOFOLLOW,
				Payload: lib.Response{Code: 200, Message: "ok", Data: p},
			}, []int{credentials.FollowerID})
		}

	} else {
		if ok {
			utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
				Action:  utils.NOTIFICATIOFOLLOW,
				Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
			}, []int{credentials.FollowerID})
		}
	}

	lib.ResponseFormatter(w, response)
}

func (f Follow) Follow(r *http.Request, response *lib.Response, credentials *models.Follow) {
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	follow := lib.CreateData{
		ID:               0,
		Credentials:      *credentials,
		Table:            "follower",
		ForeignFields:    []string{"user"},
		LookingForFields: []string{},
	}
	follow.CreateWithoutValidator(response, sqlService)
	credentials.ID = follow.ID

}

func (f Follow) UnFollow(r *http.Request, response *lib.Response, credentials models.Follow) {
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf("DELETE FROM follower WHERE follower_id = %d AND following_id = %d", credentials.FollowerID, credentials.FollowingID)
	lib.ErrorWriter(response, "ok", 200)

	if err := sqlService.Delete(query); err != nil {
		message, statusCode := lib.SqlError(err, []string{"user"}, []string{})
		lib.ErrorWriter(response, message, statusCode)
		return
	}
}

func (f Follow) AcceptARequestOrReject(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.Follow

	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.FollowingID = lib.Convert(params.Get("user_id"))
	credentials.FollowerID = userID
	credentials.Status = params.Get("status")

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil || credentials.Status == "pending" {
		message := fmt.Sprintf("%s * Invalid value. The Status field cannot contain pending value.", validators.GetValidatorErrors(errValidator))
		lib.ErrorWriter(&response, message, http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(`UPDATE follower SET status = '%s' WHERE following_id = %d AND follower_id = %d;`, credentials.Status, credentials.FollowingID, credentials.FollowerID)
	if _, insertErr := sqlService.Update(query); insertErr != nil {
		message, statusCode := lib.SqlError(insertErr, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		return
	}

	notifications := notifications.Notifications{}
	notifications.CleanNotification(r, &response, credentials.FollowingID, credentials.FollowerID, "follow_request", 0, 0)
	if response.Code != http.StatusOK {
		lib.ResponseFormatter(w, response)
		return
	}

	for _, userid := range []int{credentials.FollowerID, credentials.FollowingID} {
		client, ok := f.App.Hub.Clients[userid]

		if !ok {
			continue
		}

		payload, err := utils.GetPayloadUser(r, *f.App.Hub, userid, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATELISTUSERS",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userid})
	}

	lib.ResponseFormatter(w, response)
}
