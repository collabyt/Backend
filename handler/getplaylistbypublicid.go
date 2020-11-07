package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// GetPlaylistByPublicID returns a given playlist by it's public ID.
func GetPlaylistByPublicID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, publicID)
	if !playlist.IsPublic {
		WriteErrorReply(
			w,
			http.StatusForbidden,
		)
		return
	}
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(playlist)
	w.Write(jsonResponse)
}
