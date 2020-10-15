package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
	"github.com/gorilla/mux"
)

// DeleteVideo this is the handler that takes care of the hability to delete a
// given video from a specific playlist
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	PublicID, _ := vars["pid"]
	playlist, err := model.GetPlaylistByPublicID(database.DB, PublicID)
	if err != nil {
		errorStdTreatment(
			fmt.Errorf("Coud not locate the playlist by it's public ID"),
			w,
			http.StatusNotFound,
		)
		return
	}
	VideoID, ok := vars["vid"]
	if !ok {
		errorStdTreatment(
			fmt.Errorf("Could not identify a valid video ID"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	var v model.Video
	v.PlaylistID = playlist.ID
	v.ID, err = strconv.Atoi(VideoID)
	if err != nil {
		errorStdTreatment(
			fmt.Errorf("Video ID invalid"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	ok = model.DeleteVideo(database.DB, v)
	if !ok {
		errorStdTreatment(
			fmt.Errorf("Could not delete the requested video"),
			w,
			http.StatusInternalServerError,
		)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID), http.StatusSeeOther)
	return
}
