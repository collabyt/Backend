package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// GetKeywords return a list of (max 10) keywords to be used as part of a new
// playlist
func GetKeywords(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method GET Endpoint:/api/v1/keywords from Client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	addressParams := r.URL.Query()
	if len(addressParams["likewise"]) < 1 || len(addressParams["likewise"][0]) < 2 {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	wordList, err := model.GetKeywordsByPartialWord(database.Db, addressParams["likewise"][0])
	if err != nil {
		WriteErrorReply(w, http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(wordList)
	w.Write(jsonResponse)
}
