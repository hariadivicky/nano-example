package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

// OpenDatabase is functions to open database connectivity
func OpenDatabase(dsn string) {
	var err error
	db, err = sql.Open("sqlite3", dsn)

	handleError("could not open database: %v", err)

	if err := db.Ping(); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
}

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
