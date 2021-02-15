package api

import (
	"github.com/collabyt/Backend/collabyt"
	"github.com/labstack/echo/v4"
)

// Session holds the service for session to be used by the controllers
type Session struct {
	sessionService *collabyt.Session
}

// Routes insert into the echo general group the session related endpoints
func (s *Session) Routes(api *echo.Group) {
	api.POST("", s.requestPlaylistAccess)
	api.DELETE(":PublicID", s.revokePlaylistAccess)
}
