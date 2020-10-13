package model

import (
	"database/sql"
	"fmt"
)

// Video :
// Refers to a video, which is a part of a playlist
type Video struct {
	ID         int    `json:"id,omitempty"`
	PlaylistID int    `json:"playlistid,omitempty"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	UniqueID   string `json:"uniqueid,omitempty"`
}

// CreateVideoInPlaylist :
// Creates a video in the database including it's relation with the playlist to
// which the video belongs
func CreateVideoInPlaylist(db *sql.DB, v Video) (Video, bool) {
	row := db.QueryRow(
		`INSERT INTO public.video
		(name, link, unique_id, playlist_id)
		VALUES($1, $2, $3, $4)
		RETURNING id;`,
		v.Name,
		v.Link,
		v.UniqueID,
		v.PlaylistID,
	)
	err := row.Scan(&v.ID)
	fmt.Println(err)
	return v, err == nil
}

// DeleteVideo :
// This function deletes a video from the database. For it to work, it must
// be part of a specific playlist. Returns ok if file was found and deleted
// without problems.
func DeleteVideo(db *sql.DB, v Video) bool { //TODO

	return true
}
