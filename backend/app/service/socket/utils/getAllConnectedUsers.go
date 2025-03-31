package utils

import (
	"strconv"

	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
)

// return all users connected to the websocket server.

func GetAllConnectedUsers(hub socket.Hub) []string {
	var connectedUsers []string
	for _, singleClient := range hub.Clients {
		value := strconv.Itoa(singleClient.User.ID)
		connectedUsers = append(connectedUsers, value)
	}
	return connectedUsers
}
