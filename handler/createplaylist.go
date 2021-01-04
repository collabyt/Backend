package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// CreatePlaylist Insert a new playlist in the database based on the data
// delivered in JSON format.
func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method POST Endpoint:/api/v1/playlists from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	var playlist model.Playlist
	err := json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	playlist, err = model.CreatePlaylist(playlist)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	jsonPlaylist, err := json.Marshal(playlist)
	if err != nil {
		WriteErrorReply(w, http.StatusInternalServerError)
		return
	}
	w.Write(jsonPlaylist)
}
