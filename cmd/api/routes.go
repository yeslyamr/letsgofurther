package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)
	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods("GET")

	moviesSubrouter := router.PathPrefix("/v1/movies").Subrouter()

	moviesSubrouter.HandleFunc("/", app.listMoviesHandler).Methods("GET")
	moviesSubrouter.HandleFunc("/", app.createMovieHandler).Methods("POST")
	moviesSubrouter.HandleFunc("/{id}", app.showMovieHandler).Methods("GET")
	moviesSubrouter.HandleFunc("/{id}", app.updateMovieHandler).Methods("PATCH")
	moviesSubrouter.HandleFunc("/{id}", app.deleteMovieHandler).Methods("DELETE")
	moviesSubrouter.Use(app.requireActivatedUser)

	router.HandleFunc("/v1/users", app.registerUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/activated", app.activateUserHandler).Methods("PUT")
	router.HandleFunc("/v1/users/authentication", app.createAuthenticationTokenHandler).Methods("POST")

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
