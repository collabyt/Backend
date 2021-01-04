package handler

import (
	"fmt"

	"github.com/collabyt/Backend/model"
)

func fetchPlaylist(publicID string) (model.Playlist, error) {
	ps, err := model.GetPlaylistByPublicID(publicID)
	if err != nil {
		return model.Playlist{}, fmt.Errorf("Could not find playlist")
	}
	return ps, nil
}
