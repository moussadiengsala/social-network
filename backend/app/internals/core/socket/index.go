package socket

import (
	"time"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

const (
	WriteWait      = 10 * time.Second    // This constant represents the maximum time allowed for a write to the WebSocket connection to complete. In my case, it's set to 10 seconds.
	PongWait       = 60 * time.Second    //This constant represents the maximum time allowed for a pong message (response to a ping) to be received. It's set to 60 seconds.
	PingPeriod     = (PongWait * 9) / 10 // This constant defines the interval between sending ping messages to the WebSocket client. It's set to 90% of pongWait.
	MaxMessageSize = 1124                // This constant defines the maximum size of a message that can be received from the WebSocket client. It's set to 512 bytes.
)

type SocketEventReceiver struct {
	Action  string `json:"action"`
	Payload any    `json:"payload"`
}

type SocketEventResponse struct {
	Action  string       `json:"action"`
	Payload lib.Response `json:"payload"`
}

func setSocketPayloadReadConfig(c *Client) {
	c.WebSocketConnection.SetReadLimit(MaxMessageSize)
	c.WebSocketConnection.SetReadDeadline(time.Now().Add(PongWait))
	c.WebSocketConnection.SetPongHandler(func(string) error { c.WebSocketConnection.SetReadDeadline(time.Now().Add(PongWait)); return nil })
}

func UnRegisterAndCloseConnection(c *Client) {
	c.Hub.Unactive <- c
	c.WebSocketConnection.Close()
}
