package groups

import (
	"fmt"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/notifications"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

func (g Groups) CreateEvents(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := models.GroupEvents{}

	if err := lib.ParseForm(&credentials, r); err != nil {
		lib.ErrorWriter(&response, "Something went wrong! Make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.AuthorID = userID

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	if isAuthorUserMember := lib.IsMembersExistAndAccepted(r, credentials.GroupID, userID); !isAuthorUserMember {
		lib.ErrorWriter(&response, "the group is not found or you not part of it to perform this operation!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	events := lib.CreateData{
		ID:            0,
		Credentials:   credentials,
		Table:         "groupEvents",
		ForeignFields: []string{"user", "group"},
	}

	events.CreateWithoutValidator(&response, sqlService)
	credentials.ID = events.ID

	if response.Code != 200 {
		lib.ResponseFormatter(w, response)
		return
	}

	members, err := lib.GetAllGroupsMembers(r, credentials.GroupID, credentials.AuthorID, false)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"member"}, []string{"group", "user"})
		lib.ErrorWriter(&response, message, statusCode)
		lib.ResponseFormatter(w, response)
		return
	}

	// notify the members of the group the events
	for _, memberID := range members {
		notification := notifications.Notifications{}
		credentialsNotification := models.Notifications{
			SenderID:         credentials.AuthorID,
			ReceiverID:       memberID,
			GroupID:          credentials.GroupID,
			EventsID:         credentials.ID,
			NotificationType: "groups_events",
		}

		// Store the notification in the database
		notification.CreateNotification(r, &response, &credentialsNotification)
		if response.Code != http.StatusOK {
			lib.ResponseFormatter(w, response)
			return
		}

		client, ok := g.App.Hub.Clients[memberID]
		if !ok {
			continue
		}

		payload, err := utils.GetPayload(r, *g.App.Hub, memberID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}
		payload.Items = credentialsNotification

		// Send the notification to the following user
		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  utils.NOTIFICATIOGROUPEVENTS,
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{memberID})
	}

	lib.ResponseFormatter(w, response)
}

func (g Groups) GetAllGroupsEvents(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	credentials := []models.GetGroupEvents{}
	userID := r.Context().Value(models.UserIDKey).(int)

	params := r.URL.Query()
	groupID := lib.Convert(params.Get("group_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	events := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupEvents{},
		Query:           fmt.Sprintf(internals.QUERY_GETTING_All_EVENTS_GROUPS, userID, userID, groupID, limit, offset),
	}

	events.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetByIDGroupsEvents(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	groupID := lib.Convert(params.Get("group_id"))
	eventID := lib.Convert(params.Get("event_id"))

	credentials := models.GetGroupEvents{}
	groups := lib.GetFeed{
		Credentials: &credentials,
		Query:       fmt.Sprintf(internals.QUERYGETTINGEVENTSGROUPS, userID, userID, groupID, eventID),
	}

	groups.GetSingleFeed(r, &response)
	lib.ResponseFormatter(w, response)
}
