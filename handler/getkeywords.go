package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// GetKeywords :
// Return a list of (max 10) keywords to be used as part of a new playlist
func GetKeywords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	addressParams := r.URL.Query()
	if len(addressParams["likewise"]) < 1 || len(addressParams["likewise"][0]) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(model.Error{Description: "likewise string with at least two characters is mandatory"})
		w.Write(errRet)
		return
	}
	wordList, err := model.GetKeywordsByPartialWord(database.DB, addressParams["likewise"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		return
	}
	jsonResponse, _ := json.Marshal(wordList)
	w.Write(jsonResponse)
}
