package collabyt

type playlistRepository interface {
}

type Playlist struct {
	repo playlistRepository
}

func NewPlaylist(r playlistRepository) *Playlist {
	return &Playlist{r}
}
