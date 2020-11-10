package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// CreateKeyword insert a new keyword to the database, if it already exists,
// returns the existing one.
func CreateKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var word model.Keyword
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	if len(word.Word) < 2 {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	word, err = model.CreateKeyword(database.DB, word.Word)
	if err != nil {
		WriteErrorReply(w, http.StatusConflict)
		return
	}
	jsonResponse, _ := json.Marshal(word)
	w.Write(jsonResponse)
}
