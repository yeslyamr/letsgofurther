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

	router.Handle("/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler)).Methods("GET")
	router.Handle("/v1/movies", app.requirePermission("movies:write", app.createMovieHandler)).Methods("POST")
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler)).Methods("GET")
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler)).Methods("PATCH")
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:write", app.deleteMovieHandler)).Methods("DELETE")

	router.HandleFunc("/v1/users", app.registerUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/activated", app.activateUserHandler).Methods("PUT")
	router.HandleFunc("/v1/users/authentication", app.createAuthenticationTokenHandler).Methods("POST")

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
