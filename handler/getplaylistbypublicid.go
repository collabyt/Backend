package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// GetPlaylistByPublicID returns a given playlist by it's public ID.
func GetPlaylistByPublicID(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method GET Endpoint:/api/v1/playlists/{PublicID} from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
	}
	playlist, err := model.GetPlaylistByPublicID(publicID)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	if !playlist.IsPublic {
		WriteErrorReply(w, http.StatusUnauthorized)
		return
	}
	playlist.Passphrase = ""
	jsonResponse, _ := json.Marshal(playlist)
	w.Write(jsonResponse)
}
