package handler

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/collabyt/Backend/model"
)

func generateNewSession(size int, id int) (model.Session, error) {
	// TODO: Implement new session created log
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
