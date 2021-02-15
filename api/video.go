package api

import (
	"github.com/collabyt/Backend/collabyt"
	"github.com/labstack/echo/v4"
)

// Video holds the service for video to be used by the controllers
type Video struct {
	videoService *collabyt.Video
}

// Routes insert into the echo general group the video related endpoints
func (v *Video) Routes(api *echo.Group) {
	api.POST("/:PublicID/videos", v.insertVideoInPlaylist)
	api.DELETE("/:PublicID/videos/:VideoID", v.deassociateVideo)
}
