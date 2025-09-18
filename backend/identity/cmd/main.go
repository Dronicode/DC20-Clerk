package main

import (
	"log"
	"net/http"

	"dc20clerk/backend/identity/internal/middleware"
	"dc20clerk/backend/identity/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// Load JWKS once at startup (or periodically if needed)
	jwksProvider := middleware.NewJWKSProvider("https://your-supabase-project-url/auth/v1/.well-known/jwks.json")
	r := router.NewRouter(jwksProvider)

	log.Println("Identity service running on :8081")
	http.ListenAndServe(":8081", r)
}
