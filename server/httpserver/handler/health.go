package handler

import (
	"net/http"

	"github.com/Chad-Glazier/controlward/riot"
)

// Writes a 200 response to indicate that the server is functioning.
func Health(w http.ResponseWriter, r *http.Request) {
	err := riot.NewClient().Ping()
	if err != nil {
		http.Error(
			w, 
			"failed to connect to riot server", 
			http.StatusInternalServerError,
		)
	}
}
