package postgres

import (
	"database/sql"
	"fmt"

	// _ postgres drive
	_ "github.com/lib/pq"
)

// NewDB opens a connection to the postgres database using the environment
// variables.
func NewDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DbHost,
		DbPort,
		DbUser,
		DbPassword,
		DbDatabase,
		DbSsl,
	)
	db, err := sql.Open(DbSource, dbInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
