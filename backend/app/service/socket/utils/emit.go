package utils

import (
	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
)

func EmitToSpecificClient(hub *socket.Hub, client *socket.Client, payload socket.SocketEventResponse, userID []int) {
	for _, userid := range userID {
		targetClient, ok := hub.Clients[userid]
		if !ok {
			continue
		}

		select {
		case targetClient.Send <- payload:
		default:
			close(targetClient.Send)
			delete(hub.Clients, userid)
		}
	}
}
