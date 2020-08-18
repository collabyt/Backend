package model

// Keyword represent a single keyword database object with his own ID.
type Keyword struct {
	ID   int    `json:"id"`
	Word string `json:"keyword"`
}

// Keywords represent a list of keyword objects with their respective ID's.
type Keywords struct {
	Words []Keyword
}

// CreateKeyword creates a new keyword in the database if it doesn't exist already.
// If exists, returns the word itself from the database.
func CreateKeyword(word string) (Keyword, error) {
	var k Keyword
	return k, nil
}

// GetKeywordByWord get a keyword based on the an existing word.
func GetKeywordByWord(word string) (Keyword, error) {
	var k Keyword
	return k, nil
}

// GetKeywordsByPartialWord returns a list of keyword with a maximum amount of 10.
func GetKeywordsByPartialWord(partialWord string) (Keywords, error) {
	var ks Keywords
	return ks, nil
}
