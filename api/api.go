// Package api orgqanize and validate the endpoints and does basic request
// validation, to then call the collabyt (service) providing the necessary data
// for the business logic to proceed.
package api

import (
	"fmt"

	"github.com/collabyt/Backend/collabyt"
	"github.com/labstack/echo/v4"
)

// API holds the groups of endpoints for the application
type API struct {
	Auths     *collabyt.Auth
	Errors    *collabyt.Error
	Keywords  *collabyt.Keyword
	Playlists *collabyt.Playlist
	Sessions  *collabyt.Session
	Videos    *collabyt.Video
}

// Routes insert into echo que groups of endpoints for the application
func (a *API) Routes(e *echo.Echo) {
	g := "/api/v2"
	keyword := &Keyword{a.Keywords}
	playlist := &Playlist{a.Playlists}
	session := &Session{a.Sessions}
	video := &Video{a.Videos}
	keyword.Routes(e.Group(fmt.Sprintf("%s/keywords", g)))
	session.Routes(e.Group(fmt.Sprintf("%s/sessions", g)))
	playlist.Routes(e.Group(fmt.Sprintf("%s/playlists", g)))
	video.Routes(e.Group(fmt.Sprintf("%s/playlists", g)))
}
