package client

import "time"

type Client struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	IsOnline bool   `json:"is_online"`
}

type Message struct {
	SenderId string
	RoomId   string
	Content  string
	Date     time.Time
}
