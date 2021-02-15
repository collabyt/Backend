package repo

import "database/sql"

type Video struct {
	DB *sql.DB
}

func NewVideo(db *sql.DB) *Video {
	return &Video{db}
}
