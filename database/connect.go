package database

import (
	"database/sql"
	"fmt"

	// _ postgres drive
	_ "github.com/lib/pq"
)

var (
	// DB is the Database connection pool
	DB *sql.DB
)

// Connect opens a connection to the postgres database using the environment
// variables.
func Connect() {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbDatabase,
		dbSsl,
	)
	db, err := sql.Open(dbSource, dbInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
}
