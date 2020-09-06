package handler

import (
	"fmt"
	"net/http"
)

// LatestPlaylists :
// Return a list of the next 10 public playlists from newest to oldest, starting from the lastid parameter
// TODO LatestPlaylists: This is just a test for getting the parameters from the URL
func LatestPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	addressParams := r.URL.Query()
	if len(addressParams["keyword"]) < 1 && len(addressParams["afterid"]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if addressParams["keyword"] != nil {
		fmt.Fprintf(w, addressParams["keyword"][0])
		return
	}
	if addressParams["afterid"] != nil {
		fmt.Fprintf(w, addressParams["afterid"][0])
		return
	}
}
