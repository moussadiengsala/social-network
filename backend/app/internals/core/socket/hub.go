package socket

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	Clients                   map[int]*Client
	Active                    chan *Client
	Unactive                  chan *Client
	HandleUseractiveEvent     func(hub *Hub, client *Client)
	HandleUserDisconnectEvent func(hub *Hub, client *Client)
}

// NewHub will will give an instance of an Hub
func NewHub() *Hub {
	return &Hub{
		Active:   make(chan *Client),
		Unactive: make(chan *Client),
		Clients:  make(map[int]*Client),
	}
}

// Run will execute Go Routines to check incoming Socket events
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Active:
			hub.HandleUseractiveEvent(hub, client)
		case client := <-hub.Unactive:
			hub.HandleUserDisconnectEvent(hub, client)
		}
	}
}
