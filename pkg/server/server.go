package server

import (
	"html/template"
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
	s.router.HandleFunc("/", IndexHandler)
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	index := `
		<html>
			<head>
				<style>
				</style>
			</head>
			<body>
					<form action="/watchtower/" method="post">
						Ethereum Address:
						<input type="text" name="address">

						Email:
						<input type="text" name="email">
						
						Phone number:
						<input type="text" name="phone">

						<input type="submit" value="Submit">
					</form> 
			</body>
		</html>
	`

	t := template.New("index")
	t, err := t.Parse(index)
	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}
