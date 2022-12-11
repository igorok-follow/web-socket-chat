package client

import (
	"chat/hub"
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

type Req struct {
	RecipientId string `json:"recipient_id"`
	SenderId    string `json:"sender_id"`
	Content     string `json:"content"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Uid  string `json:"uid"`
	Send chan *Message
	Hub  *hub.Hub
	Conn *websocket.Conn
}

type Message struct {
	SenderId    string
	RecipientId string
	Content     string
	Date        time.Time
}

func ServeWs(w http.ResponseWriter, r *http.Request, h *hub.Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var req *Req
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
	}

	h.Upgrade(conn, req.SenderId)
}

func (c *Client) WritePump() {
	for {
		select {
		case m := <-c.Send:
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			body, err := json.Marshal(m)
			if err != nil {
				log.Println(err)
			}

			w.Write(body)
		}
	}
}

func (c *Client) ReadPump() {
	//defer func() {
	//c.Hub.unregister <- c
	//c.Conn.Close()
	//}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var message *Message
		err = json.Unmarshal(m, &message)
		if err != nil {
			log.Println(err)
		}

		c.Hub.Broadcast <- message
	}
}
