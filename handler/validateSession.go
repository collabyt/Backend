package handler

import (
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/model"
)

func validateSession(req *http.Request, p model.Playlist) (bool, error) {
	cook, err := req.Cookie(p.PublicID)
	if err != nil {
		return false, fmt.Errorf("Access denied to the requested playlist")
	}
	s, err := model.GetSessionBySessionID(cook.Value)
	if err != nil {
		return false, fmt.Errorf("Access denied to the requested playlist")
	}
	if s.PlaylistID != p.ID {
		return false, fmt.Errorf("Your session appears to be valid, but not for this Playlist")
	}
	return true, nil
}
