package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/jwt"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
)

type Client struct {
	Hub                       *Hub
	WebSocketConnection       *websocket.Conn
	Send                      chan SocketEventResponse
	User                      models.GetUser
	W                         http.ResponseWriter
	R                         *http.Request
	HandleSocketPayloadEvents func(client *Client, socketEventPayload SocketEventReceiver, emitIo []int)
}

func (c *Client) ReadPump() {
	var socketEventPayload SocketEventReceiver

	defer UnRegisterAndCloseConnection(c)
	setSocketPayloadReadConfig(c)

	for {
		// response := lib.Response{Code: 200, Message: "ok"}
		_, payload, err := c.WebSocketConnection.ReadMessage()
		if decoderErr := json.NewDecoder(bytes.NewReader(payload)).Decode(&socketEventPayload); decoderErr != nil {
			log.Printf("error: %v", decoderErr)
			break
		}

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error ===: %v", err)
			}
			break
		}

		if socketEventPayload.Action != "join" {
			c.HandleSocketPayloadEvents(c, socketEventPayload, []int{})
		} else {

			if err := c.handleJoinAction(socketEventPayload); err != nil {
				fmt.Println(err)
				// c.sendErrorPayload(err)
				return
			}
		}
	}
}

func (c *Client) WritePump() {
	// Create a ticker that ticks at the specified interval
	ticker := time.NewTicker(PingPeriod)

	// Ensure that the ticker is stopped and the WebSocket connection is closed when the function exits
	defer func() {
		ticker.Stop()
		c.WebSocketConnection.Close()
	}()

	for {
		select {
		case payload, ok := <-c.Send:
			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(payload)
			finalPayload := reqBodyBytes.Bytes()
			c.WebSocketConnection.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				c.WebSocketConnection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.WebSocketConnection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(finalPayload)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.WebSocketConnection.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.WebSocketConnection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleJoinAction(socketEventPayload SocketEventReceiver) error {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials, ok := socketEventPayload.Payload.(map[string]interface{})
	if !ok {
		response.Code = http.StatusInternalServerError
		response.Message = "error parsing payload"
		return fmt.Errorf("error parsing payload: %v", socketEventPayload)
	}

	token, ok := credentials["token"].(string)
	if !ok {
		response.Code = http.StatusInternalServerError
		response.Message = "error getting token"
		return fmt.Errorf("error getting token: %v", socketEventPayload)
	}

	j := jwt.JWT{}
	payload, err := j.CheckingToken(c.R, &response, token)
	if err != nil {
		return err
	}

	if response.Code == http.StatusOK {
		c.User = payload.User
		c.Hub.Active <- c
	}

	return nil
}

func (c *Client) sendErrorPayload(err error) {
	payloadError := SocketEventResponse{
		Action:  "error",
		Payload: lib.Response{Code: http.StatusInternalServerError, Message: err.Error()},
	}

	select {
	case c.Send <- payloadError:
		UnRegisterAndCloseConnection(c)
	default:
		close(c.Send)
		UnRegisterAndCloseConnection(c)
	}
}
