package router

import (
	"net/http"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/internal/handler"
	"dc20clerk/backend/identity/internal/middleware"

	"github.com/gorilla/mux"
)

// NewRouter sets up all routes and middleware.
func NewRouter(jwks *auth.JWKSProvider) http.Handler {
	r := mux.NewRouter()

	// Global middleware
	r.Use(middleware.LoggingMiddleware)

	// Auth-protected routes
	r.Handle("/profile", middleware.JWTMiddleware(jwks)(http.HandlerFunc(handler.ProfileHandler))).Methods("GET")

	// Public routes
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/register", handler.Register).Methods("POST")

	return r
}
