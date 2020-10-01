package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

//Playlist :
// represent a single playlist to be created or reproduced
type Playlist struct {
	ID         int       `json:"id"`
	PublicID   string    `json:"publicid"`
	Name       string    `json:"name"`
	IsPublic   bool      `json:"public"`
	Passphrase string    `json:"passphrase"`
	Words      []Keyword `json:"keywords"`
	Playlist   []Video   `json:"videos"`
}

// GetPlaylistByPublicID :
// Returns a single playlist to be shown in the webpage, including the video
// data.
func GetPlaylistByPublicID(db *sql.DB, publicID string) (Playlist, error) {
	pRow := db.QueryRow(
		`SELECT * FROM public.playlist WHERE public_id = $1;`,
		publicID,
	)
	var p Playlist
	err := pRow.Scan(&p)
	if err != nil {
		return Playlist{}, err
	}
	p.Playlist, err = GetVideosByPlaylistID(db, p.ID)
	if err != nil {
		return p, err
	}
	return p, nil
}

// CreatePlaylist :
// Create a new playlist, either public or private
func CreatePlaylist(db *sql.DB, playlist Playlist) (Playlist, error) {
	pass, err := bcrypt.GenerateFromPassword(
		[]byte(playlist.Passphrase),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return Playlist{}, err
	}
	playlist.Passphrase = string(pass)
	size := 6
	randomBytes := make([]byte, size)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return Playlist{}, err
	}
	playlist.PublicID = base64.URLEncoding.EncodeToString(randomBytes)
	row := db.QueryRow(
		`INSERT INTO public.playlist
		(public_id, "name", is_public, passphrase)
		VALUES($1, $2, $3, $4)
		RETURNING id;`,
		playlist.PublicID,
		playlist.Name,
		playlist.IsPublic,
		playlist.Passphrase,
	)
	err = row.Scan(&playlist.ID)
	if err != nil {
		return Playlist{}, err
	}
	err = CreateKeywordsRelation(db, playlist.ID, playlist.Words)
	if err != nil {
		return Playlist{}, err
	}
	err = CreateVideosFromPlaylist(db, playlist.ID, playlist.Playlist)
	if err != nil {
		return Playlist{}, err
	}
	playlist.Passphrase = "SECRET"
	return playlist, nil
}
