package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gw "github.com/gorilla/websocket"
	"log"
	"net/http"
)

// WebSocket ...
type WebSocket struct {
	Upgrader *gw.Upgrader
	Response http.ResponseWriter
	Request  *http.Request
}

var upgrader = gw.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
}

// NewWebSocket ...
func NewWebSocket(w http.ResponseWriter, r *http.Request) *WebSocket {
	return &WebSocket{
		Upgrader: &upgrader,
		Response: w,
		Request:  r,
	}
}

func (ws *WebSocket) ConnectionWebSocket() error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := ws.Upgrader.Upgrade(ws.Response, ws.Request, nil)
	if err != nil {
		return err
	}

	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	err = conn.WriteMessage(msgType, []byte("i readed this: "+string(msg)))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", wsHandler())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Authorization", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// Регистрация pprof-обработчиков

	log.Println("Server started")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func wsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws := NewWebSocket(w, r)
		err := ws.ConnectionWebSocket()
		if err != nil {
			log.Println(err)
		}
	})
}
