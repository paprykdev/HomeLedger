package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/paprykdev/homeledger/internal/handlers"
	"github.com/paprykdev/homeledger/internal/middleware"
)

func RegisterUserRoutes(r chi.Router, handler *handlers.UserHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})
}

func RegisterAccountRoutes(r chi.Router, handler *handlers.AccountHandler) {
	r.Route("/accounts", func(r chi.Router) {
		// All account routes require authentication
		r.Use(middleware.JWTAuth)

		r.Post("/", handler.CreateAccount)
		r.Get("/", handler.GetAccounts)
		r.Get("/{id}", handler.GetAccountById)
		r.Delete("/{id}", handler.DeleteAccount)
	})
}
