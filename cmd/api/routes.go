package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/status", app.statusHandler)
	router.GET("/movie/:id", app.getOneMovie)

	return router
}
