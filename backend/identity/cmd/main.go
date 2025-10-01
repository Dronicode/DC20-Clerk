package main

import (
	"log"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/internal/router"
	"dc20clerk/backend/identity/pkg/util"
)

func main() {
	// Load JWKS once at startup (or periodically if needed)
	jwksProvider := auth.NewJWKSProvider(util.Env("SUPABASE_URL") + "auth/v1/.well-known/jwks.json")
	jwksProvider.StartAutoRefresh(30 * time.Minute)
	r := router.NewRouter(jwksProvider)

	log.Println("Identity service running on :8081")
	http.ListenAndServe(":8081", r)

}
