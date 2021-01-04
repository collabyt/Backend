package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// CreateVideoInPlaylist Insert a video in a given playlist if the user has
// access to do so
func CreateVideoInPlaylist(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method POST Endpoint:/api/v1/playlists/{PublicID}/videos from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	playlist, err := fetchPlaylist(publicID)
	if err != nil {
		WriteErrorReply(w, http.StatusNotFound)
		return
	}
	if !playlist.IsPublic {
		ok, err := validateSession(r, playlist)
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
	video, ok := model.CreateVideoInPlaylist(video)
	if !ok {
		WriteErrorReply(w, http.StatusInternalServerError)
		return
	}
	np, err := model.GetPlaylistByPublicID(playlist.PublicID)
	np.Passphrase = ""
	jsonResponse, _ := json.Marshal(np)
	w.Write(jsonResponse)
}
