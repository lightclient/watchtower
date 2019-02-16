package main

import (
	"fmt"
	"log"

	root "github.com/c-o-l-o-r/watchtower/pkg"
	"github.com/c-o-l-o-r/watchtower/pkg/config"
	"github.com/c-o-l-o-r/watchtower/pkg/kubernetes"
	"github.com/c-o-l-o-r/watchtower/pkg/server"
)

type App struct {
	server *server.Server
	client *kubernetes.Client
	config *root.Config
}

func (a *App) Initialize() {
	a.config = config.GetConfig()

	var err error
	a.client, err = kubernetes.NewClient(a.config.Kubernetes)
	if err != nil {
		log.Fatalln("unable to connect to kubernetes")
	}

	w := kubernetes.NewWatchtowerService(a.client)

	a.server = server.NewServer(w, a.config)
}

func (a *App) Run() {
	fmt.Println("Run")
	a.server.Start()
}
