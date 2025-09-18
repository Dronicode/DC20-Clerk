package main

import (
	"log"
	"net/http"

	"dc20clerk/backend/identity/internal/middleware"
	"dc20clerk/backend/identity/internal/router"
	"dc20clerk/backend/identity/pkg/utilities"
)

func main() {
	// Load JWKS once at startup (or periodically if needed)
	jwksProvider := middleware.NewJWKSProvider(utilities.Env("SUPABASE_URL") + "auth/v1/.well-known/jwks.json")
	r := router.NewRouter(jwksProvider)

	log.Println("Identity service running on :8081")
	http.ListenAndServe(":8081", r)
}
