package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
	"github.com/gorilla/mux"
)

// GetPlaylistByPublicID returns a given playlist by it's public ID.
func GetPlaylistByPublicID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	if vars["pid"] == "" {
		errorStdTreatment(
			fmt.Errorf("Could not find the playlist public ID"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	PublicID, ok := vars["pid"]
	if !ok {
		errorStdTreatment(
			fmt.Errorf("Could not find the playlist public ID"),
			w,
			http.StatusBadRequest,
		)
		return
	}
	playlist, err := model.GetPlaylistByPublicID(database.DB, PublicID)
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
		return
	}
	jsonResponse, _ := json.Marshal(playlist)
	w.Write(jsonResponse)
}
