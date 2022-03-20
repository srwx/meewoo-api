package main

import (
	"encoding/json"
	"net/http"
)

// Convert struct to json and write json to response
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	res, err := json.MarshalIndent(wrapper, "", "\t")

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)

	return nil
}

// Write json error message to response
func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type structError struct {
		Message string `json:"message"`
	}

	isError := structError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, isError, "error")
}
