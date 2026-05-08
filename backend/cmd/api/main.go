package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/paprykdev/homeledger/internal/database"
	"github.com/paprykdev/homeledger/internal/handlers"
	"github.com/paprykdev/homeledger/internal/middleware"
	"github.com/paprykdev/homeledger/internal/routes"
)

func main() {
	db := database.New()
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	transactionHandler := handlers.NewTransactionHandler(db)

	routes.RegisterHealthRoutes(r)
	routes.RegisterTransactionRoutes(r, transactionHandler)

	log.Println("server running on :8080")

	err := chi.Walk(r, func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {

		log.Printf("[%s] %s", method, route)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", r)
}
