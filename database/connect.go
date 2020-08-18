package database

import (
	"database/sql"
	"fmt"

	"github.com/IvebeenDotIo/Backend/config/postgres"

	// _ postgres drive
	_ "github.com/lib/pq"
)

var (
	// DB is the Database connection pool
	DB *sql.DB
)

// Connect to the postgres database using the environment variables.
func Connect() {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		postgres.DbHost, postgres.DbPort, postgres.DbUser, postgres.DbPassword, postgres.DbDatabase, postgres.DbSsl)
	db, err := sql.Open(postgres.DbSource, dbInfo)
	if err != nil {

		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
}
