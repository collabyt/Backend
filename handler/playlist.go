package handler

import "net/http"

// CreateOrGetPlaylist :
// Return a playlist given it's id (playlist parameter). It returns an empty
// playlist if given playlist is protected, unless already allowed by a session
// cookie or a valid password is sent.
func CreateOrGetPlaylist(w http.ResponseWriter, r *http.Request) {

}
