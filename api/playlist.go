package api

import (
	"github.com/collabyt/Backend/collabyt"
	"github.com/labstack/echo/v4"
)

// Playlist holds the service for playlits to be used by the controllers
type Playlist struct {
	playlistService *collabyt.Playlist
}

// Routes insert into the echo general group the playlist related endpoints
func (p *Playlist) Routes(api *echo.Group) {
	api.POST("", p.createPlaylist)
	api.GET("/:PublicID", p.getPlaylistByPublicID)
	api.GET("", p.getPublicPlaylists)
}
