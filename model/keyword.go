package model

import (
	"database/sql"
)

// Keyword :
// represent a single keyword database object with his own ID.
type Keyword struct {
	ID   int    `json:"id"`
	Word string `json:"keyword"`
}

// CreateKeyword :
// creates a new keyword in the database if it doesn't exist already.
// If exists, returns the word itself from the database.
func CreateKeyword(db *sql.DB, word string) (Keyword, error) {
	var k Keyword
	sRow := db.QueryRow(`SELECT * FROM keyword WHERE word = $1;`, word)
	err := sRow.Scan(&k.ID, &k.Word)
	if err == nil {
		return k, nil
	}
	iRow := db.QueryRow(
		`INSERT INTO public.keyword 
			(word) 
		VALUES
			($1)
		RETURNING id, word;`,
		word,
	)
	err = iRow.Scan(&k.ID, &k.Word)
	if err != nil {
		return Keyword{}, err
	}
	return k, nil
}

// GetKeywordByID :
// returns a single keyword based in it's id.
func GetKeywordByID(db *sql.DB, id int) (Keyword, error) {
	row := db.QueryRow(`SELECT id, word FROM keyword WHERE id = $1;`, id)
	var k Keyword
	if err := row.Scan(&k); err != nil {
		return Keyword{}, err
	}
	return k, nil
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
