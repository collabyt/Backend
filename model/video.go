package model

// Video :
// Refers to a video, which is a part of a playlist
type Video struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	UniqueID   string `json:"uniqueid"`
	PlaylistID int    `json:"playlistid"`
}
