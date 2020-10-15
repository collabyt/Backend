package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

//Playlist represent a single playlist to be created or reproduced
type Playlist struct {
	ID         int       `json:"id,omitempty"`
	PublicID   string    `json:"publicid,omitempty"`
	Name       string    `json:"name,omitempty"`
	IsPublic   bool      `json:"public,omitempty"`
	Passphrase string    `json:"passphrase,omitempty"`
	Words      []Keyword `json:"keywords"`
	Playlist   []Video   `json:"videos"`
}

// GetPlaylistByPublicID returns a single playlist including all it's videos
// and keywords.
func GetPlaylistByPublicID(db *sql.DB, publicID string) (Playlist, error) {
	pRow := db.QueryRow(
		`SELECT 
			id,
			public_id,
			name,
			is_public
		FROM 
			public.playlist 
		WHERE 
			public_id = $1;`,
		publicID,
	)
	var p Playlist
	err := pRow.Scan(&p.ID, &p.PublicID, &p.Name, &p.IsPublic)
	if err != nil {
		return Playlist{}, err
	}
	p.Playlist, err = GetVideosByPlaylistID(db, p.ID)
	if err != nil {
		return p, err
	}
	p.Words, err = GetKeywordsByPlaylistID(db, p.ID)
	if err != nil {
		return p, err
	}
	return p, nil
}

// CreatePlaylist create a new playlist, either public or private
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
