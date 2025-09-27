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
			log.Println("[AUTH] AuthMiddleware triggered")

			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			log.Printf("[AUTH] AuthHeader: %s", authHeader)
			if !strings.HasPrefix(authHeader, "Bearer ") {
				log.Println("[AUTH] Missing or malformed header")
				http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
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
			sub, okSub := claims["sub"].(string)
			if !ok || !okSub || sub == "" {
				http.Error(w, "missing sub claim", http.StatusUnauthorized)
				return
			}

			log.Println("[AUTH] Authenticated user ID:", sub)
			// Inject user ID into context
			ctx := context.WithValue(r.Context(), UserIDKey, sub)
			ctx = context.WithValue(ctx, AccessTokenKey, tokenString)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
