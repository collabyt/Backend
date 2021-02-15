package main

import (
	"github.com/collabyt/Backend/api"
	"github.com/collabyt/Backend/collabyt"
	"github.com/collabyt/Backend/postgres"
	"github.com/collabyt/Backend/repo"
	"github.com/labstack/echo/v4"
)

func main() {
	db := postgres.NewDB()
	api := &api.API{
		Auths:     collabyt.NewAuth(repo.NewAuth(db)),
		Errors:    collabyt.NewError(repo.NewError(db)),
		Keywords:  collabyt.NewKeyword(repo.NewKeyword(db)),
		Playlists: collabyt.NewPlaylist(repo.NewPlaylist(db)),
		Sessions:  collabyt.NewSession(repo.NewSession(db)),
		Videos:    collabyt.NewVideo(repo.NewVideo(db)),
	}

	e := echo.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
