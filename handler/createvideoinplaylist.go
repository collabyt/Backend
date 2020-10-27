package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// CreateVideoInPlaylist Insert a video in a given playlist if the user has
// access to do so
func CreateVideoInPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "pid")
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
	}
	var video model.Video
	err = json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, publicID)
	if err != nil {
		errorStdTreatment(
			fmt.Errorf("Could not find a playlist with that ID"),
			w,
			http.StatusNotFound,
		)
		return
	}
	video.PlaylistID = playlist.ID
	video, ok := model.CreateVideoInPlaylist(database.DB, video)
	if !ok {
		errorStdTreatment(
			fmt.Errorf("Something wrong happened when adding the video to the Playlist"),
			w,
			http.StatusInternalServerError,
		)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID), http.StatusSeeOther)
	return
}
