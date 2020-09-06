package model

import (
	"database/sql"
	"fmt"
)

// CreateKeywordsRelation :
// Execute one single SQL instruction relating existing keywords from the
// database with a unique Playlist.
func CreateKeywordsRelation(db *sql.DB, playlistID int, wordList []Keyword) error {
	formattedInserts := "INSERT INTO public.playlist_keyword (playlist_id, keyword_id) VALUES "
	for _, word := range wordList {
		formattedInserts += fmt.Sprintf("(%d,%d),", playlistID, word.ID)
	}
	formattedInserts = formattedInserts[:len(formattedInserts)-1]
	SQLStatement, err := db.Prepare(formattedInserts)
	if err != nil {
		return err
	}
	_, err = SQLStatement.Exec()
	if err != nil {
		return err
	}
	return nil
}
