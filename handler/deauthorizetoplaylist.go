package handler

import (
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// DeauthorizeToPlaylist set the cookie related to the given playlist to expire
func DeauthorizeToPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	playlistID, err := fetchVars(r, "PublicID")
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	cook, err := r.Cookie(playlistID)
	if err != nil {
		return
	}
	cook.MaxAge = -1
	http.SetCookie(w, cook)
	err = model.DeleteSessionBySessionID(database.DB, cook.Value)
	if err != nil {
		errorStdTreatment(err, w, http.StatusInternalServerError)
		return
	}
}