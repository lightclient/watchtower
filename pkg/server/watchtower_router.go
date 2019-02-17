package server

import (
	"encoding/json"
	"net/http"

	root "github.com/c-o-l-o-r/watchtower/pkg"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type watchtowerRouter struct {
	watchtowerService root.WatchtowerService
}

func NewWatchtowerRouter(w root.WatchtowerService, router *mux.Router) *mux.Router {
	watchtowerRouter := watchtowerRouter{w}

	router.HandleFunc("/", watchtowerRouter.createWatchtowerJSONHandler).
		HeadersRegexp("Content-Type", "application/(text|json)").
		Methods("POST")

	router.HandleFunc("/", watchtowerRouter.createWatchtowerFormHandler).
		HeadersRegexp("Content-Type", "application/x-www-form-urlencoded").
		Methods("POST")

	return router
}

func (wt *watchtowerRouter) createWatchtowerJSONHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var watchtowerAttributes root.WatchtowerAttributes
	err := decoder.Decode(&watchtowerAttributes)
	if err != nil {
		panic(err)
	}

	err = wt.watchtowerService.CreateWatchtower(watchtowerAttributes)
	if err != nil {
		panic(err)
	}

	Json(w, http.StatusOK, "success")
}

func (wt *watchtowerRouter) createWatchtowerFormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	decoder := schema.NewDecoder()
	var watchtowerAttributes root.WatchtowerAttributes
	err = decoder.Decode(&watchtowerAttributes, r.PostForm)
	if err != nil {
		panic(err)
	}

	err = wt.watchtowerService.CreateWatchtower(watchtowerAttributes)
	if err != nil {
		panic(err)
	}

	Json(w, http.StatusOK, "success")
}
