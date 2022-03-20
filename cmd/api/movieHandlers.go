package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"meewoo-api/models"

	"github.com/julienschmidt/httprouter"
)

// GET /movie/:id
func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
	}

	movie := models.Movie{
		ID:          id,
		Title:       "Some title",
		Description: "Some description",
		Year:        2022,
		ReleaseDate: time.Date(2000, 8, 24, 13, 52, 49, 0, time.Local),
		Runtime:     120,
		Rating:      4,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err != nil {
		log.Println(err)
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}
