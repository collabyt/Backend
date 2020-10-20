package model

// Auth type holds the information given by the user to authenticate into a
// specific playlist.
type Auth struct {
	PlaylistID int    `json:"playlistid,omitempty"`
	PublicID   string `json:"publicid,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
}
