package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// Keyword :
// insert a new keyword to the database, if it already exists, returns the existing one.
func Keyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		return
	}
	var word model.Keyword
	if err = json.Unmarshal(body, &word); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		return
	}
	word, err = model.CreateKeyword(database.DB, word.Word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		return
	}
	jsonResponse, _ := json.Marshal(word)
	w.Write(jsonResponse)
}
