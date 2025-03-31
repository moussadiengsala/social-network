package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

// payload
type ResponseWSUser struct {
	Users []models.GetUser `json:"users"`
	Items any
}

type ResponseWSGroup struct {
	Groups []models.GetGroups `json:"groups"`
	Items  any
}

type ResponseWSUnreadNotification struct {
	UnreadNotifications int `json:"unread_notifications"`
	UnreadMessages      int `json:"unread_messages"`
	Items               any
}

type ResponseWS struct {
	Users               []models.GetUser   `json:"users"`
	Groups              []models.GetGroups `json:"groups"`
	UnreadNotifications int                `json:"unread_notifications"`
	UnreadMessages      int                `json:"unread_messages"`
	Items               any
}

func GetPayload(r *http.Request, hub socket.Hub, userID int, response *lib.Response) (ResponseWS, error) {

	var users, groups string
	var unreadNotifications int
	var result ResponseWS

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(internals.QUERY_ON_WS_PAYLOAD, strings.Join(GetAllConnectedUsers(hub), ","))

	if err := sqlService.SelectSingle(query, []interface{}{userID, userID, userID, userID, userID, userID, userID, userID, userID}, &users, &groups, &unreadNotifications); err != nil {
		message, code := lib.SqlError(err, []string{"user", "groups"}, []string{"user", "groups"})
		lib.ErrorWriter(response, message, code)
		return result, err
	}

	if err := json.Unmarshal([]byte(users), &result.Users); err != nil {
		fmt.Println(err)
		lib.ErrorWriter(response, fmt.Sprintf("Error unmarshalling users: %v", err), http.StatusInternalServerError)
		return result, err
	}

	if err := json.Unmarshal([]byte(groups), &result.Groups); err != nil {
		fmt.Println(err)
		lib.ErrorWriter(response, fmt.Sprintf("Error unmarshalling groups: %v", err), http.StatusInternalServerError)
		return result, err
	}

	result.UnreadNotifications = unreadNotifications
	return result, nil
}

func GetPayloadUser(r *http.Request, hub socket.Hub, userID int, response *lib.Response) (ResponseWSUser, error) {

	var users string
	var result ResponseWSUser

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(internals.QUERY_ON_WS_PAYLOAD_USERS, userID, userID, userID, strings.Join(GetAllConnectedUsers(hub), ","), userID, userID, userID)

	if err := sqlService.SelectSingle(query, []interface{}{}, &users); err != nil {
		message, code := lib.SqlError(err, []string{"user", "groups"}, []string{"user", "groups"})
		lib.ErrorWriter(response, message, code)
		return result, err
	}

	if err := json.Unmarshal([]byte(users), &result.Users); err != nil {
		lib.ErrorWriter(response, fmt.Sprintf("Error unmarshalling users: %v", err), http.StatusInternalServerError)
		return result, err
	}

	return result, nil
}

func GetPayloadGroups(r *http.Request, hub socket.Hub, userID int, response *lib.Response) (ResponseWSGroup, error) {

	var groups string
	var result ResponseWSGroup

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(internals.QUERY_ON_WS_PAYLOAD_GROUPS, userID, userID)

	if err := sqlService.SelectSingle(query, []interface{}{}, &groups); err != nil {
		message, code := lib.SqlError(err, []string{"user", "groups"}, []string{"user", "groups"})
		lib.ErrorWriter(response, message, code)
		return result, err
	}

	if err := json.Unmarshal([]byte(groups), &result.Groups); err != nil {
		lib.ErrorWriter(response, fmt.Sprintf("Error unmarshalling groups: %v", err), http.StatusInternalServerError)
		return result, err
	}

	return result, nil
}

func GetPayloadUnreadNotification(r *http.Request, hub socket.Hub, userID int, response *lib.Response) (ResponseWSUnreadNotification, error) {

	var unreadNotification ResponseWSUnreadNotification

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	query := fmt.Sprintf(internals.QUERY_ON_WS_PAYLOAD_NOTIFICATIONS, userID)

	if err := sqlService.SelectSingle(query, []interface{}{}, &unreadNotification.UnreadNotifications); err != nil {
		message, code := lib.SqlError(err, []string{"user", "groups"}, []string{"user", "groups"})
		lib.ErrorWriter(response, message, code)
		return unreadNotification, err
	}

	return unreadNotification, nil
}
