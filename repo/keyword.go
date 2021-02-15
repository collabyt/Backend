package repo

import "database/sql"

type Keyword struct {
	DB *sql.DB
}

func NewKeyword(db *sql.DB) *Keyword {
	return &Keyword{db}
}
