package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// DeleteVideo this is the handler that takes care of the hability to delete a
// given video from a specific playlist
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method DELETE Endpoint:/api/v1/playlists/{PublicID}/videos/{VideoID} from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
	}
	playlist, err := model.GetPlaylistByPublicID(publicID)
	if err != nil {
		WriteErrorReply(w, http.StatusNotFound)
		return
	}
	videoID, err := fetchVars(r, "VideoID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
	}
	var v model.Video
	v.PlaylistID = playlist.ID
	v.ID, err = strconv.Atoi(videoID)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	ok := model.DeleteVideo(v)
	if !ok {
		WriteErrorReply(w, http.StatusInternalServerError)
		return
	}
	http.Redirect(
		w,
		r,
		fmt.Sprintf("/api/v1/playlists/%s", playlist.PublicID),
		http.StatusSeeOther,
	)
	return
}
