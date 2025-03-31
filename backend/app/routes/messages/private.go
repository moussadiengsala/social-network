package message

import (
	"fmt"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

func (pm Message) CreatePrivateMessage(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := models.PrivateMessage{}

	if err := lib.ParseForm(&credentials, r); err != nil {
		lib.ErrorWriter(&response, "Something went wrong! Make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.SenderID = userID

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	if isFreinds := lib.IsTwoUsersAreFreinds(r, credentials.SenderID, credentials.ReceiverID); !isFreinds {
		lib.ErrorWriter(&response, "you need to be freinds with this user to perform thia operation.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	message := lib.CreateData{
		ID:            0,
		Credentials:   credentials,
		Table:         "privateMessage",
		ForeignFields: []string{"user"},
	}

	message.CreateWithoutValidator(&response, sqlService)
	credentials.ID = message.ID

	for i, userid := range []int{credentials.SenderID, credentials.ReceiverID} {
		action := []string{utils.PRIVATEMESSAGE, utils.NOTIFICATIONPRIVATEMESSAGE}
		client, ok := pm.App.Hub.Clients[userid]

		if ok {
			payload, err := utils.GetPayloadUser(r, *pm.App.Hub, userid, &response)
			if err != nil {
				message, statusCode := lib.SqlError(err, []string{"user"}, []string{})
				lib.ErrorWriter(&response, message, statusCode)
				utils.HandleUnexpectedEvent(client, &response, []int{userID})
				return
			}
			payload.Items = credentials

			utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
				Action:  action[i],
				Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
			}, []int{userid})
		}
	}

	lib.ResponseFormatter(w, response)
}

func (pm Message) GetByIDPrivateMessage(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	messageID := lib.Convert(r.URL.Query().Get("message_id"))
	userID := r.Context().Value(models.UserIDKey).(int)
	data := models.GetPrivateMessage{}

	user := lib.GetFeed{
		Credentials: &data,
		Query:       fmt.Sprintf(internals.QUERY_GETTING_PRIVATE_MESSAGE, messageID),
	}
	user.GetSingleFeed(r, &response)

	// mark the meaage as readed
	query := fmt.Sprintf(`UPDATE privateMessage SET status = "read" WHERE id = %d AND status = "unread";`, messageID)
	pm.MarkAsReadedPrivateMessage(r, &response, query)

	client, ok := pm.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadUser(r, *pm.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user"}, []string{})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATEUSERLIST",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userID})
	}

	lib.ResponseFormatter(w, response)
}

func (pm Message) GetALLPrivateMessages(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	target := lib.Convert(params.Get("target_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	credentials := []models.GetPrivateMessage{}

	var query = fmt.Sprintf(internals.QUERYGETTINGALLPRIVATEMESSAGE, userID, target, target, userID, limit, offset)
	message := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetPrivateMessage{},
		Query:           query,
	}

	message.GetAllFeed(r, &response)

	for _, m := range credentials {
		if m.ReceiverID != userID {
			continue
		}

		query := fmt.Sprintf(`UPDATE privateMessage SET status = "read" WHERE id = %d AND status = "unread";`, m.ID)
		pm.MarkAsReadedPrivateMessage(r, &response, query)
	}

	client, ok := pm.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadUser(r, *pm.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user"}, []string{})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATEUSERLIST",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userID})
	}

	lib.ResponseFormatter(w, response)
}
