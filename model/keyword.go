package model

import (
	"github.com/collabyt/Backend/database"
)

// Keyword represent a single keyword database object with his own ID.
type Keyword struct {
	ID   int    `json:"id"`
	Word string `json:"keyword"`
}

// CreateKeyword creates a new keyword in the database if it doesn't exist
//already. If exists, returns the word itself from the database.
func CreateKeyword(word string) (Keyword, error) {
	var k Keyword
	iRow := database.Db.QueryRow(
		`INSERT INTO public.keyword 
			(word) 
		VALUES
			($1)
		RETURNING id, word;`,
		word,
	)
	err := iRow.Scan(&k.ID, &k.Word)
	return k, err
}

// GetKeywordByID returns a single keyword based in it's id.
func GetKeywordByID(id int) (Keyword, error) {
	row := database.Db.QueryRow(`SELECT id, word FROM keyword WHERE id = $1;`, id)
	var k Keyword
	if err := row.Scan(&k); err != nil {
		return Keyword{}, err
	}
	return k, nil
}
