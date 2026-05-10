package main

import (
	"log"

	"github.com/paprykdev/homeledger/internal/database"
)

func main() {
	db := database.New()

	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal(err)
	}
	log.Println("migration successful")
}
