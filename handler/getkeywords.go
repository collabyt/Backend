package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// GetKeywords return a list of (max 10) keywords to be used as part of a new
// playlist
func GetKeywords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	addressParams := r.URL.Query()
	if len(addressParams["likewise"]) < 1 || len(addressParams["likewise"][0]) < 2 {
		errorStdTreatment(
			fmt.Errorf("likewise string with at least two characters is mandatory"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	wordList, err := model.GetKeywordsByPartialWord(database.DB, addressParams["likewise"][0])
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(wordList)
	w.Write(jsonResponse)
}
