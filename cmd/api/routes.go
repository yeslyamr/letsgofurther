package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)
	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.Methods(http.MethodOptions)

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods(http.MethodGet)

	router.Handle("/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler)).Methods(http.MethodGet)
	router.Handle("/v1/movies", app.requirePermission("movies:write", app.createMovieHandler)).Methods(http.MethodPost)
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler)).Methods(http.MethodGet)
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler)).Methods(http.MethodPatch)
	router.Handle("/v1/movies/{id}", app.requirePermission("movies:write", app.deleteMovieHandler)).Methods(http.MethodDelete)

	router.HandleFunc("/v1/users", app.registerUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/v1/users/activated", app.activateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/v1/users/authentication", app.createAuthenticationTokenHandler).Methods(http.MethodPost)

	router.Use(app.recoverPanic)
	router.Use(app.enableCORS)
	router.Use(app.rateLimit)
	router.Use(app.authenticate)

	return router
}
