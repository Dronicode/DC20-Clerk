package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"dc20clerk/backend/identity/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware verifies JWT and injects user ID into context.
// It expects a Bearer token in the Authorization header.
func JWTMiddleware(jwksProvider *auth.JWKSProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[AUTH] → %s %s", r.Method, r.URL.Path)

			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				log.Printf("[AUTH] ✖ Missing or malformed Authorization header")
				http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate token
			token, err := auth.ValidateToken(tokenString, jwksProvider.Get())
			if err != nil {
				log.Printf("[AUTH] ✖ Token validation failed: %v", err)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Extract subject claim
			claims, ok := token.Claims.(jwt.MapClaims)
			sub, okSub := claims["sub"].(string)
			if !ok || !okSub || sub == "" {
				log.Printf("[AUTH] ✖ Missing sub claim in token")
				http.Error(w, "missing sub claim", http.StatusUnauthorized)
				return
			}

			// Inject user ID into context
			log.Printf("[AUTH] ← Authenticated user ID: %s", sub)
			ctx := context.WithValue(r.Context(), UserIDKey, sub)
			ctx = context.WithValue(ctx, AccessTokenKey, tokenString)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
