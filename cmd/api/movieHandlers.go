package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GET /movie/:id
func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		app.errorJSON(w, err)
		app.logger.Println(errors.New("invalid id parameter"))
		return
	}

	movie, err := app.models.DB.GetOneMovie(id)

	if err != nil {
		app.logger.Println(err)
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

// GET /movies
func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := app.models.DB.GetAllMovies()
	if err != nil {
		app.logger.Panicln(err)
	}

	app.writeJSON(w, http.StatusOK, movies, "movies")
}

// GET /genres
func (app *application) getGenres(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	genres, err := app.models.DB.GetGenres()
	if err != nil {
		app.logger.Panicln(err)
	}

	app.writeJSON(w, http.StatusOK, genres, "genres")
}
