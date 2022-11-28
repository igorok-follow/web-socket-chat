package room

import "chat/client"

type Room struct {
	messageHistory []*client.Message
}
