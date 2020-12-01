package model

import (
	"fmt"

	"github.com/collabyt/Backend/database"
)

// Video refers to a video, which is a part of a playlist
type Video struct {
	ID         int    `json:"id,omitempty"`
	PlaylistID int    `json:"playlistid,omitempty"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	UniqueID   string `json:"uniqueid,omitempty"`
}

// CreateVideoInPlaylist Creates a video in the database including it's relation
// with the playlist to which the video belongs.
func CreateVideoInPlaylist(v Video) (Video, bool) {
	row := database.Db.QueryRow(
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

// DeleteVideo deletes a video from the database. For it to work, it must
// be part of a specific playlist. Returns ok if entry was found and deleted
// without problems.
func DeleteVideo(v Video) bool {
	_, err := database.Db.Exec(
		`DELETE FROM public.video
		WHERE id = $1 AND playlist_id = $2;`,
		v.ID,
		v.PlaylistID,
	)
	return err == nil
}
