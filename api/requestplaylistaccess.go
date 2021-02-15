package api

import (
	"net/http"

	"github.com/collabyt/Backend/rendering"
	"github.com/labstack/echo/v4"
)

func (s *Session) requestPlaylistAccess(ctx echo.Context) error {
	resp := rendering.LoginResponse{Success: true}
	var newAuth rendering.Auth
	err := ctx.Bind(&newAuth)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			rendering.LoginResponse{
				Success: false,
				Reason:  err.Error(),
			},
		)
	}
	s, lr := s.sessionService.Login(ctx.Request().Context(), &newAuth)
	if lr != nil {
		return ctx.JSON(
			lr.Code,
			lr,
		)
	}
	return ctx.JSON(http.StatusOK, resp)
}
