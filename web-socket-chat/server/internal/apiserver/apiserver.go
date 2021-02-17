package apiserver

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// server ...
type server struct {
	config *Config
	router *mux.Router
}

// NewServer ...
func newServer(config *Config) *server {
	return &server{
		config: config,
		router: mux.NewRouter(),
	}
}

// Run ...
func Run(config *Config) error {
	server := newServer(config)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	return http.ListenAndServe(server.config.ServerPort, handlers.CORS(headers, methods, origins)(server.router))
}

func (server *server) routers() {
	// server.router.Handle("/api/login", handlers.HandlerLoginUser()).Methods("POST")
}
