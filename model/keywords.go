package model

import "database/sql"

// Keywords :
// represent a list of keyword objects with their respective ID's.
type Keywords struct {
	Words []Keyword
}

// GetKeywordsByPartialWord :
// returns a list of keyword with a maximum amount of 10.
func GetKeywordsByPartialWord(db *sql.DB, partialWord string) (Keywords, error) {
	rows, err := db.Query(`SELECT * FROM keyword WHERE word LIKE $1 ORDER BY word LIMIT 10;`, "'"+partialWord+"%'")
	if err != nil {
		return Keywords{}, err
	}
	var ks Keywords
	for rows.Next() {
		var k Keyword
		if err := rows.Scan(&k.ID, &k.Word); err != nil {
			ks.Words = append(ks.Words, k)
		}
	}
	return ks, nil
}
