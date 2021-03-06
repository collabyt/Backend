package model

import (
	"github.com/collabyt/Backend/database"
)

// Session is the structure of a valid session in the database or cookie
type Session struct {
	PlaylistID int
	SessionID  string
}

// GetSessionBySessionID get the session from their base64 string
func GetSessionBySessionID(id string) (Session, error) {
	row := database.Db.QueryRow(
		`SELECT 
			playlist_id, session_id
		FROM 
			public."session"
		WHERE 
			session_id = $1`,
		id,
	)
	var s Session
	err := row.Scan(&s.PlaylistID, &s.SessionID)
	if err != nil {
		return Session{}, err
	}
	return s, nil
}

// CreateSession stores in the database the created session for further database
// validation
func CreateSession(s Session) error {
	var err error
	_, err = database.Db.Exec(
		`INSERT INTO public."session"
			(playlist_id, session_id)
		VALUES
			($1, $2)`,
		s.PlaylistID,
		s.SessionID,
	)
	return err
}

// DeleteSessionBySessionID erase from the database the Session by it's session
// ID string.
func DeleteSessionBySessionID(sessionID string) error {
	_, err := database.Db.Exec(
		`DELETE FROM public."session"
		WHERE session_id=$1`,
		sessionID,
	)
	return err
}
