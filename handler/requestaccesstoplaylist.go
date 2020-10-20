package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// RequestAccessToPlaylist authorize or deny access to a given playlist.
func RequestAccessToPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var auth model.Auth
	vars := mux.Vars(r)
	cook, err := r.Cookie(vars["pid"])
	if err == nil {
		s, sErr := model.GetSessionBySessionID(database.DB, cook.Value)
		if sErr != nil {
			errorStdTreatment(err, w, http.StatusInternalServerError)
			return
		}
		if s.PlaylistID != vars["pid"] {
			errorStdTreatment(err, w, http.StatusUnauthorized)
			return
		}
		playlist, err := model.GetPlaylistByPublicID(database.DB, vars["pid"])
		if err != nil {
			errorStdTreatment(err, w, http.StatusBadRequest)
			return
		}
		jsonResponse, _ := json.Marshal(playlist)
		w.Write(jsonResponse)
	}
	err = json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	if auth.PublicID == "" {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	ps, err := model.GetPlaylistByPublicID(database.DB, auth.PublicID)
	if err != nil {
		errorStdTreatment(err, w, http.StatusInternalServerError)
		return
	}
	auth.PlaylistID = ps.ID
	err = bcrypt.CompareHashAndPassword([]byte(ps.Passphrase), []byte(auth.Passphrase))
	if err != nil {
		errorStdTreatment(err, w, http.StatusUnauthorized)
		return
	}
	newCook := http.Cookie{
		Name:    auth.PublicID,
		Value:   "pending", // TODO: PENDING SUBMISSION TO DATABASE OF SESSION ID! VERY IMPORTANT
		Expires: time.Now().Add(time.Hour * 720),
	}
	http.SetCookie(w, &newCook)
	playlist, err := model.GetPlaylistByPublicID(database.DB, auth.PublicID)
	if err != nil {
		errorStdTreatment(err, w, http.StatusInternalServerError)
		return
	}
	jsonResponse, _ := json.Marshal(playlist)
	w.Write(jsonResponse)
}
