package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/dog", app.getRandomDog)

	return r
}
