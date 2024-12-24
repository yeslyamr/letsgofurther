package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods("GET")
	router.HandleFunc("/v1/movies", app.createMovieHandler).Methods("POST")
	router.HandleFunc("/v1/movies/{id}", app.showMovieHandler).Methods("GET")

	return router
}
