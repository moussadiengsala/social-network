package utils

import (
	"log"

	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

// events
func HandleSocketPayloadEvents(client *socket.Client, socketEventPayload socket.SocketEventReceiver, emitTo []int) {
	switch socketEventPayload.Action {
	case JOIN:
		log.Printf("Join Event triggered")
		HandleJoin(client, socketEventPayload)
	case DISCONNECT:
		log.Printf("Disconnect Event triggered")
		HandleDicoonect(client, socketEventPayload)
	case TYPING:
		// data := getSubmitedDataFromPayload(client, socketEventPayload, []string{"target"})
		// if data == nil {
		// 	return
		// }
		// targetClient, _ := client.Hub.Clients[data["target"]]
		// EmitToSpecificClient(client.Hub, socket.SocketEventStruct{
		// 	Action: socketEventPayload.Action,
		// 	Payload: map[string]string{
		// 		"info":        fmt.Sprintf("%s %s is typing...", client.User.Firstname, client.User.Lastname),
		// 		"sender_id":   client.User.ID,
		// 		"receiver_id": targetClient.User.ID,
		// 	},
		// }, []string{targetClient.User.ID})

	case ERROR:
		log.Printf("Error Event triggered")
		response, ok := socketEventPayload.Payload.(lib.Response)
		if !ok {
			return
		}

		EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  socketEventPayload.Action,
			Payload: response,
		}, emitTo)
	}
}
