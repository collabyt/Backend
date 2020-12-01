package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// RequestAccessToPlaylist authorize or deny access to a given playlist.
func RequestAccessToPlaylist(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement hit log
	w.Header().Set("Content-Type", "application/json")
	publicID, err := fetchVars(r, "PublicID")
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
	}
	cook, err := r.Cookie(publicID)
	if err != nil {
		var auth model.Auth
		err = json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			WriteErrorReply(w, http.StatusBadRequest)
			return
		}
		newCookie, httpErrCode, err := noCookie(auth)
		if err != nil {
			WriteErrorReply(w, httpErrCode)
			return
		}
		http.SetCookie(w, &newCookie)
		playlist, _ := model.GetPlaylistByPublicID(database.Db, newCookie.Name)
		playlist.Passphrase = ""
		jasonPlaylist, _ := json.Marshal(playlist)
		w.Write(jasonPlaylist)
		return
	}
	jsonPlaylist, err := hasCookie(cook)
	if err != nil {
		WriteErrorReply(w, http.StatusUnauthorized)
	}
	// TODO: Implement access granted log
	w.Write(jsonPlaylist)
}
