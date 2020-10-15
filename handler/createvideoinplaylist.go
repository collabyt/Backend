package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
	"github.com/gorilla/mux"
)

// CreateVideoInPlaylist Insert a video in a given playlist if the user has
// access to do so
func CreateVideoInPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	PublicID, _ := vars["pid"]
	var video model.Video
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: err.Error()},
		)
		w.Write(errRet)
		return
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, PublicID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errRet, _ := json.Marshal(
			model.Error{Description: "Could not find a playlist with that ID"},
		)
		w.Write(errRet)
		return
	}
	video.PlaylistID = playlist.ID
	video, ok := model.CreateVideoInPlaylist(database.DB, video)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		errRet, _ := json.Marshal(
			model.Error{Description: "Something wrong happened when adding the video to the Playlist."},
		)
		w.Write(errRet)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID), http.StatusSeeOther)
	return
}
