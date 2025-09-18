package router

import (
	"net/http"

	"dc20clerk/backend/identity/internal/handler"
	"dc20clerk/backend/identity/internal/middleware"

	"github.com/gorilla/mux"
)

// NewRouter sets up all routes and middleware.
func NewRouter(jwks *middleware.JWKSProvider) http.Handler {
	r := mux.NewRouter()

	// Global middleware
	r.Use(middleware.LoggingMiddleware)

	// Auth-protected routes
	r.Handle("/identity/profile", middleware.AuthMiddleware(jwks)(http.HandlerFunc(handler.ProfileHandler))).Methods("GET")

	// Public routes
	r.HandleFunc("/identity/login", handler.Login).Methods("POST")
	r.HandleFunc("/identity/register", handler.Register).Methods("POST")

	return r
}
