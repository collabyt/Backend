package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// CreateVideoInPlaylist Insert a video in a given playlist if the user has
// access to do so
func CreateVideoInPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	playlist, err := fetchPlaylist(database.DB, publicID)
	if err != nil {
		WriteErrorReply(w, http.StatusNotFound)
		return
	}
	if !playlist.IsPublic {
		ok, err := validateSession(database.DB, r, playlist)
		if err != nil {
			WriteErrorReply(w, http.StatusForbidden)
			return
		}
		if !ok {
			WriteErrorReply(w, http.StatusInternalServerError)
			return
		}
	}
	var video model.Video
	err = json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	video.PlaylistID = playlist.ID
	video, ok := model.CreateVideoInPlaylist(database.DB, video)
	if !ok {
		WriteErrorReply(w, http.StatusInternalServerError)
		return
	}
	np, err := model.GetPlaylistByPublicID(database.DB, playlist.PublicID)
	np.Passphrase = ""
	jsonResponse, _ := json.Marshal(np)
	w.Write(jsonResponse)
}
