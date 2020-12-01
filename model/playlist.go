package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/collabyt/Backend/database"
	"golang.org/x/crypto/bcrypt"
)

//Playlist represent a single playlist to be created or reproduced
type Playlist struct {
	ID         int       `json:"id,omitempty"`
	PublicID   string    `json:"publicid,omitempty"`
	Name       string    `json:"name,omitempty"`
	IsPublic   bool      `json:"public,omitempty"`
	Passphrase string    `json:"passphrase,omitempty"`
	Keywords   []Keyword `json:"keywords,omitempty"`
	Videos     []Video   `json:"videos,omitempty"`
}

// GetPlaylistByPublicID returns a single playlist including all it's videos
// and keywords.
func GetPlaylistByPublicID(publicID string) (Playlist, error) {
	pRow := database.Db.QueryRow(
		`SELECT 
			id,
			public_id,
			name,
			is_public,
			passphrase
		FROM 
			public.playlist 
		WHERE 
			public_id = $1;`,
		publicID,
	)
	var p Playlist
	err := pRow.Scan(&p.ID, &p.PublicID, &p.Name, &p.IsPublic, &p.Passphrase)
	if err != nil {
		return Playlist{}, err
	}
	p.Videos, err = GetVideosByPlaylistID(p.ID)
	if err != nil {
		return p, err
	}
	p.Keywords, err = GetKeywordsByPlaylistID(p.ID)
	if err != nil {
		return p, err
	}
	return p, nil
}

// CreatePlaylist create a new playlist, either public or private
func CreatePlaylist(playlist Playlist) (Playlist, error) {
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
	row := database.Db.QueryRow(
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
	err = CreateKeywordsRelation(playlist.ID, playlist.Keywords)
	if err != nil {
		return Playlist{}, err
	}
	err = CreateVideosFromPlaylist(playlist.ID, playlist.Videos)
	if err != nil {
		return Playlist{}, err
	}
	playlist.ID = 0
	playlist.Passphrase = "SECRET"
	if !playlist.IsPublic {
		playlist.Videos = []Video{}
		playlist.Keywords = []Keyword{}
	}
	return playlist, nil
}
