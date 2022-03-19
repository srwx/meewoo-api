package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /status
func (app *application) statusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	currentStatus := AppStatus{
		Version:    version,
		Enviroment: "development",
		Status:     "OK",
	}

	res, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
