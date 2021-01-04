package model

import (
	"fmt"

	"github.com/collabyt/Backend/database"
)

// GetKeywordsByPartialWord returns a list of keyword with a maximum amount of 10.
func GetKeywordsByPartialWord(partWord string) ([]Keyword, error) {
	rows, err := database.Db.Query(
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

// GetKeywordsByID Given a slice of id's, this will return a list a keywords
// with all keyword from such slice.
func GetKeywordsByID(ids []int) ([]Keyword, error) {
	var formattedIDsList string
	for _, id := range ids {
		if formattedIDsList != "" {
			formattedIDsList += ", " + fmt.Sprint(id)
		} else {
			formattedIDsList += fmt.Sprint(id)
		}
	}
	formattedIDsList = "(" + formattedIDsList + ")"
	rows, err := database.Db.Query(
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

// GetKeywordsByPlaylistID get all keywords that are associated with a given
// Playlist
func GetKeywordsByPlaylistID(playlistID int) ([]Keyword, error) {
	kRows, err := database.Db.Query(
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
	defer kRows.Close()
	var ks []Keyword
	for kRows.Next() {
		var k Keyword
		kRows.Scan(&k.ID, &k.Word)
		ks = append(ks, k)
	}
	return ks, nil
}
