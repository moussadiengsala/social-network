package utils

import socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"

// HandleUserDisconnectEvent will handle the Disconnect event for socket users
func HandleUserDisconnectEvent(hub *socket.Hub, client *socket.Client) {
	_, ok := hub.Clients[client.User.ID]
	if ok {
		delete(hub.Clients, client.User.ID)
		close(client.Send)

		HandleSocketPayloadEvents(client, socket.SocketEventReceiver{
			Action:  DISCONNECT,
		}, []int{})
	}
}
