package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.GET("/status", app.statusHandler)
	router.GET("/movie/:id", app.getOneMovie)
	router.GET("/movies", app.getAllMovies)
	router.GET("/genres", app.getGenres)

	return app.enableCORS(router)
}
