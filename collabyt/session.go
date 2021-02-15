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

	// TODO: Login in Playlist
	// Search Playlist by public ID, then from that search, get the playlist
	// ID, then, use this playlist id to create a valid session to store in
	// the database.
	// +-------------+
	// | session     |    +-------------+
	// +-------------+    | playlist    |
	// | session_id  |    +-------------|
	// | playlist_id |----| playlist_id |
	// +-------------+    | ...         |
	//                    +-------------+
	// Now that there is a valid session with a playlist_id that for sure
	// exists, function can proceed creating the session to send to the
	// endpoint to be transformed in a cookie.
	return nil, nil
}
