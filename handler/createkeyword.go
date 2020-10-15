package handler

import (
	"encoding/json"
	"fmt"
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
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	if len(word.Word) < 2 {
		errorStdTreatment(
			fmt.Errorf("Keyword need to have at least two characters"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	word, err = model.CreateKeyword(database.DB, word.Word)
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(word)
	w.Write(jsonResponse)
}
