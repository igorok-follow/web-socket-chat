package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (server *server) HandlerLoginUser() http.Handler {
	type User struct {
		Login string `json:"login"`
	}
	var user User
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes := server.requestReader(r)
		json.Unmarshal(bytes, &user)
	})
}

func (server *server) requestReader(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	return body
}

func (server *server) responseWriter(statusCode int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(json)
	if err != nil {
		log.Panic(err)
	}
}
