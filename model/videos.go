package model

import "database/sql"

// Videos :
// Refers to a list of the object Video.
type Videos struct {
	Videos []Video `json:"videos"`
}

// GetVideosByPlaylistID :
// retrieve all videos that belong to a specific playlist using the id from the
// database.
func GetVideosByPlaylistID(db *sql.DB, playlistID int) (Videos, error) {
	vRows, err := db.Query(
		`SELECT  id, name, link, unique_id, playlist_id 
		FROM  public.video
		WHERE playlist_id = $1`,
		playlistID,
	)
	if err != nil {
		return Videos{}, err
	}
	var (
		vs Videos
		v  Video
	)
	for vRows.Next() {
		vRows.Scan(&v)
		vs.Videos = append(vs.Videos, v)
	}
	return vs, nil
}
