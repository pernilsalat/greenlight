package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedResponse)
	// Healthcheck
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	// Movies routes
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.CreateMovieHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.GetMovieHandler)

	return router
}
