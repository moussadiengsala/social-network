package utils

import (
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

func HandleUnexpectedEvent(client *socket.Client, paylaod *lib.Response, emitTo []int) {
	HandleSocketPayloadEvents(client, socket.SocketEventReceiver{
		Action:  ERROR,
		Payload: paylaod,
	}, emitTo)
}
