package message

import (
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

func (pm Message) CreateGroupMessage(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := models.GroupMessage{}

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

	if isMember := lib.IsMembersExistAndAccepted(r, credentials.GroupID, credentials.SenderID); !isMember {
		lib.ErrorWriter(&response, "the group is not found or you not part of it to perform this operation!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	// storing the message to the database
	pm.InserGroupMessage(r, &response, &credentials)
	if response.Code != http.StatusOK {
		lib.ResponseFormatter(w, response)
		return
	}

	// get the membership of the groups
	members, err := lib.GetAllGroupsMembers(r, credentials.GroupID, 0, true)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"member"}, []string{"group", "user"})
		lib.ErrorWriter(&response, message, statusCode)
		lib.ResponseFormatter(w, response)
		return
	}

	for _, memberID := range members {
		client, ok := pm.App.Hub.Clients[memberID]
		if !ok {
			continue
		}

		action := utils.NOTIFICATIONGROUPMESSAGE
		if memberID == credentials.SenderID {
			action = utils.GROUPSMESSAGE
		}

		payload, err := utils.GetPayloadGroups(r, *pm.App.Hub, memberID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"group"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}
		payload.Items = credentials

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  action,
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{memberID})
	}

	lib.ResponseFormatter(w, response)
}

func (pm Message) InserGroupMessage(r *http.Request, response *lib.Response, credentials *models.GroupMessage) {

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	c := func(response *lib.Response, id int) error {
		members, err := lib.GetAllGroupsMembers(r, credentials.GroupID, credentials.SenderID, false)
		if err != nil || len(members) == 0 {
			return err
		}

		pairs := []string{}
		for _, member := range members {
			pairs = append(pairs, fmt.Sprintf("(%d,%d)", member, id))
		}

		query := fmt.Sprintf("INSERT INTO groupMessageStatus(user_id, group_message_id) VALUES %s", strings.Join(pairs, ", "))
		if _, insertPostErr := sqlService.Create(query); insertPostErr != nil {
			return insertPostErr
		}
		return nil
	}

	message := lib.CreateData{
		ID:            0,
		Credentials:   *credentials,
		Callback:      c,
		Table:         "groupMessage",
		ForeignFields: []string{"user", "group"},
	}
	message.CreateWithoutValidator(response, sqlService)
	credentials.ID = message.ID
}

func (pm Message) GetByIDGroupsMessage(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	data := models.GetGroupMessage{}
	messageID := lib.Convert(r.URL.Query().Get("message_id"))
	userID := r.Context().Value(models.UserIDKey).(int)

	user := lib.GetFeed{
		Credentials: &data,
		Query:       fmt.Sprintf(internals.QUERY_GETTING_GROUPS_MESSAGE, messageID),
	}
	user.GetSingleFeed(r, &response)

	// mark the meaage as readed
	query := fmt.Sprintf(`UPDATE groupMessageStatus SET status = "read" WHERE user_id = %d AND group_message_id = %d AND status = "unread";`, userID, messageID)
	pm.MarkAsReadedPrivateMessage(r, &response, query)

	lib.ResponseFormatter(w, response)
	client, ok := pm.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadGroups(r, *pm.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"group"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATELISTGROUPS",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userID})
	}

}


func (pm Message) GetALLGroupsMessages(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	group := lib.Convert(params.Get("group_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	credentials := []models.GetGroupMessage{}

	var query = fmt.Sprintf(internals.QUERY_GETTING_ALL_GROUPS_MESSAGE, group, limit, offset)
	message := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupMessage{},
		Query:           query,
	}

	message.GetAllFeed(r, &response)

	for _, m := range credentials {
		query := fmt.Sprintf(`UPDATE groupMessageStatus SET status = "read" WHERE user_id = %d AND group_message_id = %d AND status = "unread";`, userID, m.ID)
		pm.MarkAsReadedPrivateMessage(r, &response, query)
	}

	client, ok := pm.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadGroups(r, *pm.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"group"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}
		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATELISTGROUPS",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{userID})
	}

	lib.ResponseFormatter(w, response)
}
