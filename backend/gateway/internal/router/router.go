package router

import (
	"dc20clerk/backend/gateway/internal/middleware"
	"dc20clerk/backend/gateway/internal/proxy"
	"dc20clerk/common/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	log.Println("[GATEWAY] → Initializing router")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	log.Println("[GATEWAY] → Logging middleware attached")

	cfg := config.LoadGatewayConfig()
	identityURL := cfg.IdentityURL

	log.Printf("[GATEWAY] → Proxying /api/identity to %s", identityURL)

	r.PathPrefix("/api/identity/").Handler(
		http.StripPrefix("/api/identity", proxy.NewReverseProxy(identityURL)),
	)

	log.Println("[GATEWAY] ← Router setup complete")
	return r
}
