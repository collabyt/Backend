package repo

import (
	"context"
	"database/sql"

	"github.com/collabyt/Backend/model"
	"github.com/collabyt/Backend/rendering"
)

type Session struct {
	DB *sql.DB
}

func NewSession(db *sql.DB) *Session {
	return &Session{db}
}

func (s *Session) Login(
	ctx context.Context,
	a *model.Auth,
) (
	*model.Session,
	*rendering.LoginResponse,
) {

}
