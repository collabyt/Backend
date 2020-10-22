package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	vars := mux.Vars(r)
	cook, err := r.Cookie(vars["pid"])
	if err != nil {
		var auth model.Auth
		err = json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			errorStdTreatment(err, w, http.StatusBadRequest)
			return
		}
		newCookie, httpErrCode, err := noCookie(auth)
		if err != nil {
			errorStdTreatment(err, w, httpErrCode)
			return
		}
		http.SetCookie(w, &newCookie)
		playlist, _ := model.GetPlaylistByPublicID(database.DB, newCookie.Name)
		playlist.Passphrase = ""
		jasonPlaylist, _ := json.Marshal(playlist)
		w.Write(jasonPlaylist)
		return
	}
	jsonPlaylist, err := hasCookie(cook)
	if err != nil {
		errorStdTreatment(err, w, http.StatusUnauthorized)
	}
	w.Write(jsonPlaylist)
}

func hasCookie(cook *http.Cookie) ([]byte, error) {
	s, err := model.GetSessionBySessionID(database.DB, cook.Value)
	if err != nil {
		return []byte{}, err
	}
	var es model.Session
	if s == es {
		cook.MaxAge = -1
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, cook.Name)
	if err != nil {
		return []byte{}, err
	}
	jp, _ := json.Marshal(playlist)
	return jp, nil
}

func noCookie(a model.Auth) (http.Cookie, int, error) {
	if a.PublicID == "" {
		return http.Cookie{},
			http.StatusUnauthorized,
			fmt.Errorf("Invalid Public ID")
	}
	ps, err := model.GetPlaylistByPublicID(database.DB, a.PublicID)
	if err != nil {
		return http.Cookie{},
			http.StatusUnauthorized, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(ps.Passphrase),
		[]byte(a.Passphrase),
	)
	if err != nil {
		return http.Cookie{},
			http.StatusUnauthorized,
			fmt.Errorf("Wrong password, access denied")
	}
	s, err := generateNewSession(12, ps.ID)
	if err != nil {
		return http.Cookie{},
			http.StatusInternalServerError,
			fmt.Errorf("Something wrong happened")
	}

	err = model.CreateSession(database.DB, s)
	if err != nil {
		return http.Cookie{},
			http.StatusInternalServerError,
			fmt.Errorf("Could not create session in the database")
	}
	nc := http.Cookie{
		Name:    a.PublicID,
		Value:   s.SessionID,
		Expires: time.Now().Add(time.Hour * 360),
	}
	return nc,
		http.StatusOK,
		nil
}

func generateNewSession(size int, id int) (model.Session, error) {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	var es model.Session
	if err != nil {
		return es, err
	}
	return model.Session{
		PlaylistID: id,
		SessionID:  base64.URLEncoding.EncodeToString(randomBytes),
	}, nil
}
