package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite", "storage/homeledger.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id TEXT PRIMARY KEY NOT NULL,
		amount REAL NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		deleted_at DATETIME
	)
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
