package main

import (
	"chat/client"
	"chat/hub"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type RegisterClientRequest struct {
	Login string
}

func main() {
	h := hub.NewHub()
	h.Run()

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	router.Handle("/register", Register(h)).Methods("POST")
	router.Handle("/ws", Ws(h))

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Authorization", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server started")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func Ws(h *hub.Hub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r, h)
	})
}

func Register(h *hub.Hub) http.Handler {
	var register *RegisterClientRequest
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &register)
		if err != nil {
			log.Println(err)
		}

		h.Register <- &client.Client{
			Uid:  register.Login,
			Send: make(chan *client.Message),
			Hub:  h,
			Conn: nil,
		}

		fmt.Fprintf(w, "ok")
	})
}
