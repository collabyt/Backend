package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// CreatePlaylist :
// Insert a new playlist in the database based on the data delivered in JSON
// format.
// TODO CreatePlaylist has not been implemented.
func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var playlist model.Playlist
	err := json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: err.Error()},
		)
		w.Write(errRet)
		return
	}
	playlist, err = model.CreatePlaylist(database.DB, playlist)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: err.Error()},
		)
		w.Write(errRet)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID), http.StatusSeeOther)
	return
}
