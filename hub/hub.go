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
