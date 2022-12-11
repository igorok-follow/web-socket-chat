package main

import (
	"chat/client"
	"chat/hub"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	MESSAGE_TO = 1
)

type WSRequest struct {
	Type int `json:"type"`
}

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
	router.Handle("/get/people", GetPeople()).Methods("POST")
	router.Handle("/register", Register()).Methods("POST")
	router.Handle("/ws", Ws(h))

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Authorization", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server started")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func Ws(h *hub.Hub) http.Handler {
	var ws *WSRequest
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &ws)
		if err != nil {
			log.Println(err)
		}

		switch ws.Type {
		case 0:
			return
		case 1:
			client.ServeWs(w, r, h)
		}
	})
}

func Register() http.Handler {
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

		// func which can show all active clients
	})
}

func GetPeople() http.Handler {
	//var client *RegisterClientRequest
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//body, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	log.Println(err)
		//}
		//err = json.Unmarshal(body, &client)
		//if err != nil {
		//	log.Println(err)
		//}

		// func which can show all active clients
	})
}
