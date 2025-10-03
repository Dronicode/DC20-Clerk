package router

import (
	"log"
	"net/http"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/internal/handler"
	"dc20clerk/backend/identity/internal/middleware"

	"github.com/gorilla/mux"
)

// NewRouter sets up all routes and middleware.
func NewRouter(jwks *auth.JWKSProvider) http.Handler {
	log.Println("[IDENTITY] → Initializing router")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	log.Println("[IDENTITY] → Logging middleware attached")

	log.Println("[IDENTITY] → Creating routes")
	// Auth-protected routes
	r.Handle("/profile", middleware.JWTMiddleware(jwks)(http.HandlerFunc(handler.ProfileHandler))).Methods("GET")
	log.Println("[IDENTITY] → Protected route: GET /profile")

	// Public routes
	r.HandleFunc("/login", handler.Login).Methods("POST")
	log.Println("[IDENTITY] → Public route: POST /login")

	r.HandleFunc("/register", handler.Register).Methods("POST")
	log.Println("[IDENTITY] → Public route: POST /register")

	log.Println("[IDENTITY] ← Router setup complete")
	return r
}
