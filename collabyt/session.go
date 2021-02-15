package collabyt

import (
	"context"
	"net/http"

	"github.com/collabyt/Backend/model"
	"github.com/collabyt/Backend/rendering"
)

type sessionRepository interface {
	Login(ctx context.Context, a *model.Auth) (*model.Session, *rendering.LoginResponse)
}

type Session struct {
	repo sessionRepository
}

func NewSession(r sessionRepository) *Session {
	return &Session{r}
}

func (s *Session) Login(
	ctx context.Context,
	a *model.Auth,
) (
	*model.Session,
	*rendering.LoginResponse,
) {
	if len(a.PublicID) != 8 {
		return &model.Session{},
			&rendering.LoginResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Reason:  "Invalid Playlist",
			}
	}
	if len(a.Passphrase) < 1 {
		return &model.Session{},
			&rendering.LoginResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Reason:  "Invalid passphrase",
			}
	}
	s, err := s.repo.Login(ctx, a)
	if err != nil {
		return &model.Session{},
			&rendering.LoginResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Reason:  "Impossible to login",
			}
	}
	return s, nil
}
