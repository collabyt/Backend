package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/collabyt/Backend/model"
	"golang.org/x/crypto/bcrypt"
)

func hasCookie(cook *http.Cookie) ([]byte, error) {
	existingSession, err := model.GetSessionBySessionID(cook.Value)
	if err != nil {
		return []byte{}, err
	}
	var emptySession model.Session
	if existingSession == emptySession {
		cook.MaxAge = -1
		return []byte{}, errors.New("expired session")
	}
	playlist, err := model.GetPlaylistByPublicID(cook.Name)
	if err != nil {
		return []byte{}, err
	}
	playlist.Passphrase = ""
	jp, _ := json.Marshal(playlist)
	return jp, nil
}

func noCookie(a model.Auth) (http.Cookie, int, error) {
	if a.PublicID == "" {
		return http.Cookie{},
			http.StatusUnauthorized,
			fmt.Errorf("Invalid Public ID")
	}
	ps, err := model.GetPlaylistByPublicID(a.PublicID)
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

	err = model.CreateSession(s)
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
