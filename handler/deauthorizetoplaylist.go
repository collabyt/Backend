package handler

import (
	"fmt"
	"net/http"
)

// DeauthorizeToPlaylist set the cookie related to the given playlist to expire
func DeauthorizeToPlaylist(w http.ResponseWriter, r *http.Request) {
	publicID, err := fetchVars(r, "pid")
	if err != nil {
		errorStdTreatment(err, w, http.StatusBadRequest)
	}
	fmt.Println(publicID)
	// NOT FINISHED!
}
