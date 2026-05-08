package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/paprykdev/homeledger/internal/handlers"
)

func RegisterTransactionRoutes(
	r chi.Router,
	handler *handlers.TransactionHandler,
) {
	r.Route("/transactions", func (r chi.Router) {
		r.Post("/", handler.CreateTransaction)
		r.Get("/", handler.GetTransactions)
		r.Get("/{id}", handler.GetTransactionById)
		r.Delete("/{id}", handler.DeleteTransaction)
	})
}

func RegisterHealthRoutes(r chi.Router) {
	r.Get("/health", handlers.HealthCheck)
}
