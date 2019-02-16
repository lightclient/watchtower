package server

import (
	"net/http"

	root "github.com/c-o-l-o-r/watchtower/pkg"
	"github.com/gorilla/mux"
)

type watchtowerRouter struct {
	watchtowerService root.WatchtowerService
}

func NewWatchtowerRouter(w root.WatchtowerService, router *mux.Router) *mux.Router {
	watchtowerRouter := watchtowerRouter{w}
	router.HandleFunc("/", watchtowerRouter.createWatchtowerHandler).Methods("POST")
	return router
}

func (ur *watchtowerRouter) createWatchtowerHandler(w http.ResponseWriter, r *http.Request) {
	Json(w, http.StatusOK, "ok")
}
