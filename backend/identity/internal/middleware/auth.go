package middleware

import (
	"context"
	"log"
	"net/http"

	"dc20clerk/backend/identity/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

// AuthMiddleware verifies JWT and injects user ID into context.
func AuthMiddleware(jwksProvider *JWKSProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("[AUTH] AuthMiddleware triggered")

			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			log.Printf("[AUTH] AuthHeader: %s", authHeader)
			if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
				log.Println("[AUTH] Missing or malformed header")
				http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenString := authHeader[7:]
			log.Println("[AUTH] Extracted token:", tokenString)

			// Validate token
			token, err := auth.ValidateToken(tokenString, jwksProvider.Get())
			if err != nil {
				log.Printf("[AUTH] Token validation failed: %v\n", err)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Extract subject claim
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("[AUTH] Failed to parse claims")
				http.Error(w, "invalid claims", http.StatusUnauthorized)
				return
			}
			sub, ok := claims["sub"].(string)
			if !ok || sub == "" {
				log.Println("[AUTH] Missing sub claim")
				http.Error(w, "missing sub claim", http.StatusUnauthorized)
				return
			}

			log.Println("[AUTH] Authenticated user ID:", sub)
			// Inject user ID into context
			ctx := context.WithValue(r.Context(), UserIDKey, sub)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
