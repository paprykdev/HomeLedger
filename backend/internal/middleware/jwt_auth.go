package middleware

import (
	"context"
	"net/http"

	"github.com/paprykdev/homeledger/internal/auth"
)

// JWTAuth middleware validates JWT tokens and adds user_id to context
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from header
		tokenString, err := auth.ExtractTokenFromRequest(r)
		if err != nil {
			http.Error(w, "unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Parse and validate token
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add user_id to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
