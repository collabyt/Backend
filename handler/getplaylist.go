package handler

import (
	"net/http"
)

// GetPlaylist :
// Get a playlist given the ID passed in the URI
// TODO GetPlaylist has not been implemented.
func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	return
}
