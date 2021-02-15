package postgres

import "os"

var (
	// DbHost is the Postgres Host
	DbHost = os.Getenv("DB_HOST")
	// DbUser is the Postgres username
	DbUser = os.Getenv("DB_USER")
	// DbPassword is the password to db
	DbPassword = os.Getenv("DB_PASSWORD")
	// DbDatabase is the name of the database
	DbDatabase = os.Getenv("DB_DATABASE")
	// DbPort is the post open in the postgres server
	DbPort = os.Getenv("DB_PORT")

	// DbSsl is the should ssl be active or not
	DbSsl = os.Getenv("DB_SSL")
	// DbSource is the Source of the database
	DbSource = os.Getenv("DB_SOURCE")
)
