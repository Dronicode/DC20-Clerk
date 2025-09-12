package main

import (
	"dc20clerk/backend/gateway/internal/proxy"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load(".env")

  r := mux.NewRouter()
  r.PathPrefix("/identity/").Handler(proxy.NewReverseProxy("http://localhost:8081"))
  r.PathPrefix("/characters/").Handler(proxy.NewReverseProxy("http://localhost:8082"))

  log.Println("Gateway running on :8080")
  http.ListenAndServe(":8080", r)
}
