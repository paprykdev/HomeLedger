package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/paprykdev/homeledger/internal/handlers"
	"github.com/paprykdev/homeledger/internal/middleware"
)

func RegisterTransactionRoutes(
	r chi.Router,
	handler *handlers.TransactionHandler,
) {
	r.Route("/transactions", func (r chi.Router) {
		// All transaction routes require authentication
		r.Use(middleware.JWTAuth)

		r.Post("/", handler.CreateTransaction)
		r.Get("/", handler.GetTransactions)
		r.Get("/{id}", handler.GetTransactionById)
		r.Patch("/{id}", handler.UpdateTransaction)
		r.Delete("/{id}", handler.DeleteTransaction)
	})
}

func RegisterHealthRoutes(r chi.Router) {
	r.Get("/health", handlers.HealthCheck)
}
