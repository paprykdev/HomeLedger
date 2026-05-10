package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/paprykdev/homeledger/internal/auth"
	"github.com/paprykdev/homeledger/internal/config"
	"github.com/paprykdev/homeledger/internal/database"
	"github.com/paprykdev/homeledger/internal/handlers"
	"github.com/paprykdev/homeledger/internal/middleware"
	"github.com/paprykdev/homeledger/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize JWT secret from config
	auth.InitJWTSecret(cfg.JWTSecret)

	// Initialize database
	db := database.New()
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)
	accountHandler := handlers.NewAccountHandler(db)
	transactionHandler := handlers.NewTransactionHandler(db)

	// Register routes
	routes.RegisterHealthRoutes(r)
	routes.RegisterUserRoutes(r, userHandler)
	routes.RegisterAccountRoutes(r, accountHandler)
	routes.RegisterTransactionRoutes(r, transactionHandler)

	log.Printf("server running on :%s", cfg.Port)

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

	http.ListenAndServe(":"+cfg.Port, r)
}
