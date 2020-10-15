package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/model"
)

// GetPlaylists returns a list of playlist based on the Limit and offset.
func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limitSlc, ok := r.URL.Query()["limit"]
	limit := 10
	if ok && len(limitSlc[0]) > 0 {
		var err error
		limit, err = strconv.Atoi(limitSlc[0])
		if err != nil {
			limit = 10
		}
	}
	if limit > 25 {
		errRet, _ := json.Marshal(model.Error{Description: "The maximum allowed limit is 25"})
		w.Write(errRet)
		return
	}
	offsetSlc, ok := r.URL.Query()["offset"]
	var offset int
	if ok && len(offsetSlc[0]) > 0 {
		var err error
		offset, err = strconv.Atoi(offsetSlc[0])
		if err != nil {
			offset = 0
		}
	}
	ps, err := model.GetPlaylistsByLimitAndOffset(database.DB, limit, offset)
	if err != nil {
		errRet, _ := json.Marshal(model.Error{Description: err.Error()})
		w.Write(errRet)
		return
	}
	jsonResponse, _ := json.Marshal(ps)
	w.Write(jsonResponse)
}
