package router

import (
	"dc20clerk/backend/gateway/internal/proxy"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/api/identity/").Handler(
		http.StripPrefix("/api/identity", proxy.NewReverseProxy(os.Getenv("IDENTITY_URL"))),
	)

	return r
}
