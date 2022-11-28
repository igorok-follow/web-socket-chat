package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gw "github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var upgrader = gw.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type RegisterClientRequest struct {
	Login string
}

type RegisterClientResponse struct {
	Id      string
	Success bool
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	router.Handle("/get/people", GetPeople()).Methods("POST")
	router.Handle("/register", nil).Methods("POST")
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveClient(w, r)
	})

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Authorization", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server started")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func serveClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// to rooms hub, contains uid, recipient_id, returning room id
}

func GetPeople() http.Handler {
	var client *RegisterClientRequest
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &client)
		if err != nil {
			log.Println(err)
		}

		// func which can show all active clients
	})
}
