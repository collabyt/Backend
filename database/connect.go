package database

import (
	"database/sql"
	"fmt"

	// _ postgres drive
	"github.com/collabyt/Backend/logger"
	_ "github.com/lib/pq"
)

var (
	// Dd is the Database connection pool
	Db *sql.DB
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
	logger.Info.Println("Trying to connect to PostgreSQL database...")
	db, err := sql.Open(dbSource, dbInfo)
	if err != nil {
		logger.Error.Println("Connection to PostgreSQL database failed!")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Error.Println("Connection to PostgreSQL database failed!")
		panic(err)
	}
	Db = db
	logger.Info.Println("Successfully connected to PostgreSQL database")
}
