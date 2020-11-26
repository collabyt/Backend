package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// CreateKeyword insert a new keyword to the database, if it already exists,
// returns the existing one.
func CreateKeyword(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method POST Endpoint:/api/v1/keywords from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	var word model.Keyword
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		logger.Warning.Printf(
			"Failed Request at Endpoint: /api/v1/keywords from Client %s: %s",
			r.RemoteAddr,
			http.StatusText(http.StatusBadRequest),
		)
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	if len(word.Word) < 2 {
		logger.Warning.Printf(
			"Failed Request at Endpoint: /api/v1/keywords from Client %s: %s",
			r.RemoteAddr,
			http.StatusText(http.StatusBadRequest),
		)
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	word, err = model.CreateKeyword(database.Db, word.Word)
	if err != nil {
		logger.Warning.Printf(
			"Failed Request at Endpoint: /api/v1/keywords from Client %s: %s",
			r.RemoteAddr,
			http.StatusText(http.StatusConflict),
		)
		WriteErrorReply(w, http.StatusConflict)
		return
	}
	jsonResponse, _ := json.Marshal(word)
	w.Write(jsonResponse)
}
