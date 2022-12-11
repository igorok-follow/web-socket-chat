package hub

import (
	"chat/client"
	"chat/room"
)

type Hub struct {
	broadcast  chan *client.Message
	register   chan *client.Client
	unregister chan *client.Client
	rooms      map[string]*room.Room
	clients    map[string]*client.Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *client.Message),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
		rooms:      make(map[string]*room.Room),
		clients:    make(map[string]*client.Client),
	}
}

func (h *Hub) Run() {
	go func() {
		for {
			select {
			case m := <-h.broadcast:
				h.rooms[m.RoomId].AddMessage(m)
			}
		}
	}()
}
