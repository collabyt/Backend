package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// GetPublicPlaylists returns a list of playlist based on the Limit and offset.
func GetPublicPlaylists(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("HIT! Method POST Endpoint:/api/v1/playlists from Client %s", r.RemoteAddr)
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
		WriteErrorReply(w, http.StatusBadRequest)
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
	ps, err := model.GetPublicPlaylistsByLimitAndOffset(database.Db, limit, offset)
	if err != nil {
		WriteErrorReply(w, http.StatusInternalServerError)
		return
	}
	jsonResponse, _ := json.Marshal(ps)
	w.Write(jsonResponse)
}
