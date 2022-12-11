package client

import (
	"chat/hub"
	gw "github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = gw.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	IsOnline bool   `json:"is_online"`
	send     chan *Message
	hub      *hub.Hub
}

type Message struct {
	SenderId string
	RoomId   string
	Content  string
	Date     time.Time
}

func ServeWs(w http.ResponseWriter, r *http.Request, h *hub.Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

}

func (c *Client) WritePump() {
	for {
		select {
		case m := <-c.send:

		}
	}
}

func (c *Client) ReadPump() {

}
