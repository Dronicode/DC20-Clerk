package router

import (
	"dc20clerk/backend/gateway/internal/proxy"
	"os"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter()

  r.PathPrefix("/identity/").Handler(proxy.NewReverseProxy(os.Getenv("IDENTITY_URL")))

  return r
}
