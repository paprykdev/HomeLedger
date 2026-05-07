package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("sqlite", "homeledger.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createTables(db)

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	log.Println("server running on :8080")

	http.ListenAndServe(":8080", r)
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
