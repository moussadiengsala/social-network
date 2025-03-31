package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

// CreateNewSocketUser creates a new socket user
func CreateNewSocketUser(w http.ResponseWriter, r *http.Request, hub *socket.Hub, connection *websocket.Conn) {
	client := &socket.Client{
		Hub:                       hub,
		WebSocketConnection:       connection,
		Send:                      make(chan socket.SocketEventResponse),
		W:                         w,
		R:                         r,
		HandleSocketPayloadEvents: utils.HandleSocketPayloadEvents,
	}

	go client.WritePump()
	go client.ReadPump()
}
