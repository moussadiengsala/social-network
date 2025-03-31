package utils

import (
	socket "learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

func HandleJoin(client *socket.Client, socketEventPayload socket.SocketEventReceiver) {
	response := lib.Response{Code: 200, Message: "ok"}

	friends, err := lib.GetUsersFriends(client.R, client.User.ID)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		HandleUnexpectedEvent(client, &response, []int{client.User.ID})
		return
	}

	paload, err := GetPayload(client.R, *client.Hub, client.User.ID, &response)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		HandleUnexpectedEvent(client, &response, []int{client.User.ID})
		return
	}

	for _, userid := range friends {
		_, ok := client.Hub.Clients[userid]
		if !ok {
			continue
		}

		friendsPayload, err := GetPayloadUser(client.R, *client.Hub, userid, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			HandleUnexpectedEvent(client, &response, []int{client.User.ID})
			return
		}

		EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  NEWJOIN,
			Payload: lib.Response{Code: 200, Message: "ok", Data: friendsPayload},
		}, []int{userid})
	}

	EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
		Action:  JOIN,
		Payload: lib.Response{Code: 200, Message: "ok", Data: paload},
	}, []int{client.User.ID})
}

func HandleDicoonect(client *socket.Client, socketEventPayload socket.SocketEventReceiver) {
	response := lib.Response{Code: 200, Message: "ok"}

	friends, err := lib.GetUsersFriends(client.R, client.User.ID)
	if err != nil {
		message, statusCode := lib.SqlError(err, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		HandleUnexpectedEvent(client, &response, []int{client.User.ID})
		return
	}

	for _, userid := range friends {
		_, ok := client.Hub.Clients[userid]
		if !ok {
			continue
		}

		friendsPayload, err := GetPayloadUser(client.R, *client.Hub, userid, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user", "follower"}, []string{"user"})
			lib.ErrorWriter(&response, message, statusCode)
			HandleUnexpectedEvent(client, &response, []int{client.User.ID})
			return
		}

		EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  NEWDISCONNECT,
			Payload: lib.Response{Code: 200, Message: "ok", Data: friendsPayload},
		}, []int{userid})
	}

}
