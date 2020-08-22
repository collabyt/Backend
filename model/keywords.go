package model

import (
	"database/sql"
)

// Keywords :
// represent a list of keyword objects with their respective ID's.
type Keywords struct {
	Words []Keyword `json:"keywords"`
}

// GetKeywordsByPartialWord :
// returns a list of keyword with a maximum amount of 10.
func GetKeywordsByPartialWord(db *sql.DB, partWord string) (Keywords, error) {
	rows, err := db.Query(
		`SELECT id, word FROM keyword 
		WHERE  word LIKE $1 
		ORDER BY word 
		LIMIT 10;`,
		partWord+"%",
	)
	if err != nil {
		return Keywords{}, err
	}
	var (
		ks Keywords
		k  Keyword
	)
	for rows.Next() {
		rows.Scan(&k)
		ks.Words = append(ks.Words, k)
	}
	return ks, nil
}
