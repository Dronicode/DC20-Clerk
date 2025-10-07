package main

import (
	"log"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/internal/router"
)

func main() {

	jwksProvider := auth.NewJWKSProvider()
	jwksProvider.StartAutoRefresh(30 * time.Minute)

	r := router.NewRouter(jwksProvider)

	log.Println("[IDENTITY] ← Service ready on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("[IDENTITY] ✖ Server failed: %v", err)
	}
}
