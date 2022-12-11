package room

import "chat/client"

type Room struct {
	messageHistory []*client.Message
}

func (r *Room) AddMessage(message *client.Message) {
	r.messageHistory = append(r.messageHistory, message)
}
