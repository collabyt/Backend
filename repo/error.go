package repo

import "database/sql"

type Error struct {
	DB *sql.DB
}

func NewError(db *sql.DB) *Error {
	return &Error{db}
}
