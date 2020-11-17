package database

import "os"

var (
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbPort     = os.Getenv("DB_PORT")

	dbSsl    = os.Getenv("DB_SSL")
	dbSource = os.Getenv("DB_SOURCE")
)
