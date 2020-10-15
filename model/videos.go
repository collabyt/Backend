package model

import (
	"database/sql"
)

// GetVideosByPlaylistID retrieve all videos that belong to a specific playlist
// using the id from the database.
func GetVideosByPlaylistID(db *sql.DB, playlistID int) ([]Video, error) {
	vRows, err := db.Query(
		`SELECT  name, link, unique_id
		FROM  public.video
		WHERE playlist_id = $1
		ORDER BY id`,
		playlistID,
	)
	if err != nil {
		return []Video{}, err
	}
	var vs []Video
	for vRows.Next() {
		var v Video
		vRows.Scan(&v.Name, &v.Link, &v.UniqueID)
		vs = append(vs, v)
	}
	return vs, nil
}

// CreateVideosFromPlaylist insert into the database all the videos inserted
// with the playlist. Returns only nil or the error received from the database.
func CreateVideosFromPlaylist(db *sql.DB, playlistID int, vs []Video) error {
	stmt, err := db.Prepare("INSERT INTO public.video(name, link, unique_id, playlist_id) VALUES( $1, $2, $3, $4 )")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range vs {
		_, err := stmt.Exec(v.Name, v.Link, v.UniqueID, playlistID)
		if err != nil {
			return err
		}
	}
	return nil
}
