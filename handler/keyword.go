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
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var word model.Keyword
	if err = json.Unmarshal(body, &word); err != nil {
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	word, err = model.CreateKeyword(database.DB, word.Word)
	if err != nil {
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(word)
	w.Write(jsonResponse)
}
