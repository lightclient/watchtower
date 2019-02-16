package server

import (
	"log"
	"net/http"
	"os"

	root "github.com/c-o-l-o-r/watchtower/pkg"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

func NewServer(w root.WatchtowerService, config *root.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: config.Server,
	}

	NewWatchtowerRouter(w, s.getSubrouter("/watchtower"))
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
