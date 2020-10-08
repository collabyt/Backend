package handler

import "net/http"

// CreateVideoInPlaylist :
// Insert a video in a given playlist if the user has access to do so
func CreateVideoInPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}
