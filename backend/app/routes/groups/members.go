package groups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/notifications"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

func (g Groups) CreateMembers(r *http.Request, response *lib.Response, credentials *models.GroupMembers) {
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	member := lib.CreateData{
		ID:               0,
		Credentials:      *credentials,
		Table:            "groupMembers",
		ForeignFields:    []string{"user", "group"},
		LookingForFields: []string{"member"},
	}

	member.CreateWithoutValidator(response, sqlService)
}

func (g Groups) Accept(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	userID := r.Context().Value(models.UserIDKey).(int)
	params := r.URL.Query()
	groupID := lib.Convert(params.Get("group_id"))
	targetID := lib.Convert(params.Get("member_id"))
	action := params.Get("action")
	resp := params.Get("response")
	var query string
	var message string
	notificationsType := ""

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	if resp != "rejected" && resp != "accepted" {
		lib.ErrorWriter(&response, "invalid action.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	switch action {
	case "invite":
		query = fmt.Sprintf(`UPDATE groupMembers SET status = "%s" WHERE group_id = %d AND user_id = %d AND status = 'invited';`, resp, groupID, userID)
		message = "You must be invited by a member of this group to perform this action."
		notificationsType = "groups_invited"
	case "request":
		admin := lib.GetGroupsAdminMember(r, groupID)
		if admin != userID || admin == 0 {
			response = lib.Response{Message: "You need to be the admin of the group to perform this operation.", Code: 400}
			lib.ResponseFormatter(w, response)
			return
		}
		query = fmt.Sprintf(`UPDATE groupMembers SET status = "%s" WHERE group_id = %d AND user_id = %d AND status = 'requested';`, resp, groupID, targetID)
		message = "this user must request to join this group to perform this action."
		notificationsType = "groups_requested"
	default:
		lib.ErrorWriter(&response, "invalid action.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	id, err := sqlService.Update(query)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"member"}, []string{})
		lib.ErrorWriter(&response, message, statusCode)
		lib.ResponseFormatter(w, response)
		return
	}

	if id == 0 {
		response = lib.Response{Code: 400, Message: message}
		lib.ResponseFormatter(w, response)
		return
	}

	notifications := notifications.Notifications{}
	notifications.CleanNotification(r, &response, targetID, userID, notificationsType, groupID, 0)
	if response.Code != http.StatusOK {
		lib.ResponseFormatter(w, response)
		return
	}

	client, ok := g.App.Hub.Clients[targetID]
	if ok {
		payload, err := utils.GetPayloadGroups(r, *g.App.Hub, targetID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{userID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  "UPDATELISTGROUPS",
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{targetID})
	}

	lib.ResponseFormatter(w, response)
}

func (g Groups) GetAllGroupsMembers(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	credentials := []models.GetGroupMembers{}
	groupID := lib.Convert(r.URL.Query().Get("group_id"))

	member := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupMembers{},
		Query:           fmt.Sprintf(internals.QUERYGETTINGAllMEMBERSGROUPS, groupID),
	}

	member.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (g Groups) RequestOrInviteJoinGroup(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := models.GroupMembers{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		lib.ErrorWriter(&response, "Something went wrong! Make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	credentials.Role = "user"
	switch credentials.Status {
	case "invited":
		g.ProcessInvitation(r, &response, &credentials)
	case "requested":
		g.ProcessRequest(r, &response, &credentials)
	default:
		lib.ErrorWriter(&response, "invalid action: you need to specify the action invited or requested", http.StatusBadRequest)

	}

	lib.ResponseFormatter(w, response)
}

func (g Groups) ProcessRequest(r *http.Request, response *lib.Response, credentials *models.GroupMembers) {
	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.User = strconv.Itoa(userID)
	admin := lib.GetGroupsAdminMember(r, credentials.GroupID)

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(*credentials); errValidator != nil {
		lib.ErrorWriter(response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		return
	}

	if isAuthorUserMember := lib.IsMembersExist(r, credentials.GroupID, userID); isAuthorUserMember {
		lib.ErrorWriter(response, "You are already the member of the group or request or invite or rejected.", http.StatusBadRequest)
		return
	}

	// create a new membership with status "requested"
	g.CreateMembers(r, response, credentials)
	if response.Code != http.StatusOK {
		return
	}
	// notify the admin groups if everything is ok
	notification := notifications.Notifications{}
	credentialsNotification := models.Notifications{
		SenderID:         userID,
		ReceiverID:       admin,
		GroupID:          credentials.GroupID,
		NotificationType: fmt.Sprintf("groups_%s", credentials.Status),
	}

	if isNotificationExist := lib.IsNotificationExist(r, credentialsNotification); !isNotificationExist {
		notification.CreateNotification(r, response, &credentialsNotification)
		if response.Code != http.StatusOK {
			return
		}

		client, ok := g.App.Hub.Clients[admin]
		if ok {
			payload, err := utils.GetPayloadUnreadNotification(r, *g.App.Hub, admin, response)
			if err != nil {
				message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
				lib.ErrorWriter(response, message, statusCode)
				utils.HandleUnexpectedEvent(client, response, []int{userID})
				return
			}
			payload.Items = credentialsNotification

			utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
				Action:  utils.NOTIFICATIOGROUPREQUESTED,
				Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
			}, []int{admin})
		}
	}
}

func (g Groups) ProcessInvitation(r *http.Request, response *lib.Response, credentials *models.GroupMembers) {

	userID := r.Context().Value(models.UserIDKey).(int)
	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(*credentials); errValidator != nil {
		lib.ErrorWriter(response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		return
	}

	usersInvited := strings.Split(credentials.User, ",")
	if lib.Contains(usersInvited, strconv.Itoa(userID)) {
		lib.ErrorWriter(response, "you cannot invite yourself.", http.StatusBadRequest)
		return
	}

	isAuthorUserMember := lib.IsMembersExistAndAccepted(r, credentials.GroupID, userID)
	for _, userid := range usersInvited {
		isInvitedUserMember := lib.IsMembersExist(r, credentials.GroupID, lib.Convert(userid))
		if !isAuthorUserMember || isInvitedUserMember {
			lib.ErrorWriter(response, "meke sure to be part of the group and the user you invited is not already a member of the group or invited to join it.", http.StatusBadRequest)
			return
		}
	}

	for _, user := range usersInvited {
		userid := lib.Convert(user)
		c := credentials
		credentials.User = user

		// create a new membership with status "requested"
		g.CreateMembers(r, response, c)
		if response.Code != http.StatusOK {
			return
		}

		// notify the admin groups if everything is ok
		notification := notifications.Notifications{}
		credentialsNotification := models.Notifications{
			SenderID:         userID,
			ReceiverID:       userid,
			GroupID:          credentials.GroupID,
			NotificationType: fmt.Sprintf("groups_%s", credentials.Status),
		}

		if isNotificationExist := lib.IsNotificationExist(r, credentialsNotification); !isNotificationExist {
			notification.CreateNotification(r, response, &credentialsNotification)
			if response.Code != http.StatusOK {
				return
			}

			client, ok := g.App.Hub.Clients[userid]
			if ok {
				payload, err := utils.GetPayloadUnreadNotification(r, *g.App.Hub, userid, response)
				if err != nil {
					message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
					lib.ErrorWriter(response, message, statusCode)
					utils.HandleUnexpectedEvent(client, response, []int{userID})
					return
				}
				payload.Items = credentialsNotification

				utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
					Action:  utils.NOTIFICATIOGROUPREQUESTED,
					Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
				}, []int{userid})
			}
		}
	}
}
