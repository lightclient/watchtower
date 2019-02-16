package server

import (
	"encoding/json"
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

func (wt *watchtowerRouter) createWatchtowerHandler(w http.ResponseWriter, r *http.Request) {
	watchtowerAttributes, err := decodeWatchtowerAttributes(r)
	if err != nil {
		panic(err)
	}

	err = wt.watchtowerService.CreateWatchtower(watchtowerAttributes)
	if err != nil {
		panic(err)
	}

	Json(w, http.StatusOK, "success")
}

func decodeWatchtowerAttributes(r *http.Request) (root.WatchtowerAttributes, error) {
	decoder := json.NewDecoder(r.Body)

	var w root.WatchtowerAttributes
	err := decoder.Decode(&w)

	return w, err
}
