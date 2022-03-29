package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) enableCORS(next *httprouter.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	}
}
