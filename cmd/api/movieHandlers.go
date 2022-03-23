package main

import (
	"errors"
	"log"
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
		log.Println(err)
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}
