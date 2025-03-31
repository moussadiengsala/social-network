package utils

import (
	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
)

// HandleUserRegisterEvent will handle the Join event for New socket users
func HandleUseractiveEvent(hub *socket.Hub, client *socket.Client) {
	hub.Clients[client.User.ID] = client

	HandleSocketPayloadEvents(client, socket.SocketEventReceiver{
		Action: JOIN,
	}, []int{})
}
