package repo

import (
	"database/sql"
)

type Playlist struct {
	DB *sql.DB
}

func NewPlaylist(db *sql.DB) *Playlist {
	return &Playlist{db}
}
