package model

// Link represent a single youtube video to be embedded
type Link struct {
	Link           string `json:"link"`
	UniqueID       string `json:"uniqueid"`
	EmbeddableLink string `json:"embeddablelink"`
}

// Links represent a list of links, aka, the proper playlist
type Links struct {
	Links []Link
}

//Playlist represent a single playlist to be created or reproduced
type Playlist struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	Public     bool     `json:"public"`
	Passphrase string   `json:"passphrase"`
	Words      Keywords `json:"keywords"`
	Links
}
