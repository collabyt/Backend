package handler

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusNotFound)
		errRet, _ := json.Marshal(
			model.Error{Description: "Coud not locate the playlist by it's public ID!"},
		)
		w.Write(errRet)
	}
	VideoID, ok := vars["vid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: "Could not identify the video ID"},
		)
		w.Write(errRet)
		return
	}
	var v model.Video
	v.PlaylistID = playlist.ID
	v.ID, err = strconv.Atoi(VideoID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: "Video ID invalid"},
		)
		w.Write(errRet)
		return
	}
	ok = model.DeleteVideo(database.DB, v)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		errRet, _ := json.Marshal(
			model.Error{Description: "Could not delete the requested video."},
		)
		w.Write(errRet)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID), http.StatusSeeOther)
	return
}
