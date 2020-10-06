package model

// Video :
// Refers to a video, which is a part of a playlist
type Video struct {
	ID         int    `json:"id,omitempty"`
	PlaylistID int    `json:"playlistid,omitempty"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	UniqueID   string `json:"uniqueid,omitempty"`
}
