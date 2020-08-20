package model

import (
	"database/sql"
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Playlist :
// represent a single playlist to be created or reproduced
type Playlist struct {
	ID         int      `json:"id"`
	PublicID   string   `json:"publicid"`
	Name       string   `json:"name"`
	IsPublic   bool     `json:"public"`
	Passphrase string   `json:"passphrase"`
	Words      Keywords `json:"keywords"`
	Playlist   Videos   `json:"videos"`
}

// GetPlaylistByPublicID :
// Returns a single playlist to be shown in the webpage, including the video data.
func GetPlaylistByPublicID(db *sql.DB, publicID string) (Playlist, error) {
	var p Playlist
	return p, nil
}

// CreatePlaylist :
// Create a new playlist, either public or private
func CreatePlaylist(db *sql.DB, playlist Playlist) (Playlist, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(playlist.Passphrase), 14)
	if err != nil {
		return Playlist{}, err
	}
	playlist.Passphrase = string(pass)
	playlist.PublicID = base64.StdEncoding.EncodeToString([]byte(playlist.Name + time.Now().String()))
	row := db.QueryRow(
		`INSERT INTO public.playlist
		(public_id, "name", is_public, passphrase)
		VALUES($1, $2, $3, $4)
		RETURNING id, public_id, name, is_public;`,
		playlist.PublicID,
		playlist.Name,
		playlist.IsPublic,
		playlist.Passphrase,
	)
	var p Playlist
	err = row.Scan(&p.ID, &p.PublicID, &p.Name, &p.IsPublic)
	if err != nil {
		return Playlist{}, err
	}
	return p, nil
}
