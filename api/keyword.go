package api

import (
	"github.com/collabyt/Backend/collabyt"
	"github.com/labstack/echo/v4"
)

// Keyword holds the service for keywords to be used by the controllers
type Keyword struct {
	keywordService *collabyt.Keyword
}

// Routes insert into the echo general group the keyword related endpoints
func (k *Keyword) Routes(api *echo.Group) {
	api.POST("", k.createKeyword)
	api.GET("", k.getKeywords)
}
