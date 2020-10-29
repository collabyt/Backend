package handler

import (
	"database/sql"
	"fmt"

	"github.com/collabyt/Backend/model"
)

func fetchPlaylist(db *sql.DB, publicID string) (model.Playlist, error) {
	ps, err := model.GetPlaylistByPublicID(db, publicID)
	if err != nil {
		return model.Playlist{}, fmt.Errorf("Could not find a playlist with that ID")
	}
	return ps, nil
}
