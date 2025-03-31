package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

// Here, an Upgrader is created with specified read and write buffer sizes.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // return true, allowing connections from any origin.
}

func Ws(app *core.App) {

	app.Hub = socket.NewHub()
	app.Hub.HandleUseractiveEvent = utils.HandleUseractiveEvent
	app.Hub.HandleUserDisconnectEvent = utils.HandleUserDisconnectEvent
	go app.Hub.Run()

	app.GET("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			fmt.Println(err)
			return
		}
		CreateNewSocketUser(w, r, app.Hub, conn)
	})
}
