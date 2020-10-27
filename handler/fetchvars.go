package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func fetchVars(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	if vars[name] == "" {
		return "", fmt.Errorf("Could not find the playlist public ID")
	}
	return vars[name], nil
}
