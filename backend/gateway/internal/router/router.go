package router

import (
	"dc20clerk/backend/gateway/internal/middleware"
	"dc20clerk/backend/gateway/internal/proxy"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	log.Println("[GATEWAY] → Initializing router")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	log.Println("[GATEWAY] → Logging middleware attached")

	identityURL := os.Getenv("IDENTITY_URL")
	log.Printf("[GATEWAY] → Proxying /api/identity to %s", identityURL)

	r.PathPrefix("/api/identity/").Handler(
		http.StripPrefix("/api/identity", proxy.NewReverseProxy(identityURL)),
	)

	log.Println("[GATEWAY] ← Router setup complete")
	return r
}
