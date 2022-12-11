package hub

import (
	"chat/client"
	"github.com/gorilla/websocket"
	"sync"
)

var lock = sync.RWMutex{}

type Hub struct {
	Broadcast  chan *client.Message
	Register   chan *client.Client
	Unregister chan *client.Client
	Clients    map[string]*client.Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *client.Message),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
		Clients:    make(map[string]*client.Client),
	}
}

func (h *Hub) Upgrade(conn *websocket.Conn, id string) {
	h.Clients[id].Conn = conn
}

func (h *Hub) Run() {
	go func() {
		for {
			select {
			case c := <-h.Register:
				func(c *client.Client) {
					lock.Lock()
					defer lock.Unlock()
					h.Clients[c.Uid] = c
				}(c)
			case c := <-h.Unregister:
				func(c *client.Client) {
					lock.Lock()
					defer lock.Unlock()
					delete(h.Clients, c.Uid)
				}(c)
			case m := <-h.Broadcast:
				if c, ok := h.Clients[m.RecipientId]; ok {
					c.Send <- m
				}
			}
		}
	}()
}
