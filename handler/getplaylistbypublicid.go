package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
	"github.com/gorilla/mux"
)

// GetPlaylistByPublicID :
// Returns a given playlist by it's public ID.
func GetPlaylistByPublicID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	if vars["pid"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: "Could not find the Playlist public ID"},
		)
		w.Write(errRet)
		return
	}
	PublicID, ok := vars["pid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: "Could not find the Playlist public ID"},
		)
		w.Write(errRet)
		return
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, PublicID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRet, _ := json.Marshal(
			model.Error{Description: err.Error()},
		)
		w.Write(errRet)
		return
	}
	jsonResponse, _ := json.Marshal(playlist)
	w.Write(jsonResponse)
}
