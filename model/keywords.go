package model

import (
	"database/sql"
)

// GetKeywordsByPartialWord :
// returns a list of keyword with a maximum amount of 10.
func GetKeywordsByPartialWord(db *sql.DB, partWord string) ([]Keyword, error) {
	rows, err := db.Query(
		`SELECT id, word FROM keyword 
		WHERE  word LIKE $1 
		ORDER BY word 
		LIMIT 10;`,
		partWord+"%",
	)
	if err != nil {
		return []Keyword{}, err
	}
	var (
		ks []Keyword
		k  Keyword
	)
	for rows.Next() {
		rows.Scan(&k.ID, &k.Word)
		ks = append(ks, k)
	}
	return ks, nil
}

// GetKeywordsByID :
// Given a slice of id's, this will return a list a keywords with all keyword
// from such slice.s
func GetKeywordsByID(db *sql.DB, ids []int) ([]Keyword, error) {
	var formattedIDsList string
	for _, id := range ids {
		if formattedIDsList != "" {
			formattedIDsList += ", " + string(id)
		} else {
			formattedIDsList += string(id)
		}
	}
	formattedIDsList = "(" + formattedIDsList + ")"
	rows, err := db.Query(
		`SELECT id, word FROM keyword WHERE id IN $1;`,
		formattedIDsList,
	)
	if err != nil {
		return []Keyword{}, err
	}
	var ks []Keyword
	for rows.Next() {
		var k Keyword
		rows.Scan(&k)
		ks = append(ks, k)
	}
	return ks, nil
}

// GetKeywordsByPlaylistID :
// Get all keywords that are associated with a given Playlist
func GetKeywordsByPlaylistID(db *sql.DB, playlistID int) ([]Keyword, error) {
	kRows, err := db.Query(
		`SELECT 
			k.id,
			k.word
		FROM 
			playlist_keyword AS pk
		INNER JOIN
			keyword AS k
		ON
			pk.keyword_id = k.id
		WHERE
			pk.playlist_id = $1;`,
		playlistID,
	)
	if err != nil {
		return []Keyword{}, err
	}
	var ks []Keyword
	for kRows.Next() {
		var k Keyword
		kRows.Scan(&k.ID, &k.Word)
		ks = append(ks, k)
	}
	return ks, nil
}
