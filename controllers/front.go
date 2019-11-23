package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers register all the handlers for the request to the application.
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", uc)
	http.Handle("/users/", uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
