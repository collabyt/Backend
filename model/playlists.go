package model

import (
	"database/sql"
)

// GetPlaylistsByLimitAndOffset :
// Given a limit and offset, returns a list of PUBLIC playlists from the database
func GetPlaylistsByLimitAndOffset(db *sql.DB, limit int, offset int) ([]Playlist, error) {
	pRows, err := db.Query(`
	SELECT 
		id, public_id, name
	FROM 
		playlist
	WHERE
		is_public = true
	ORDER BY
		id DESC
	LIMIT 
		$1
	OFFSET
		$2`,
		limit,
		offset,
	)
	if err != nil {
		return []Playlist{}, err
	}
	var ps []Playlist
	for pRows.Next() {
		var p Playlist
		err = pRows.Scan(&p.ID, &p.PublicID, &p.Name)
		if err != nil {
			return []Playlist{}, err
		}
		p.Words, err = GetKeywordsByPlaylistID(db, p.ID)
		if err != nil {
			return []Playlist{}, err
		}
		p.Playlist, err = GetVideosByPlaylistID(db, p.ID)
		if err != nil {
			return []Playlist{}, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
