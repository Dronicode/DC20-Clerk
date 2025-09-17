package main

import (
	"dc20clerk/backend/identity/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}

func main() {
  godotenv.Load(".env")

  r := mux.NewRouter()
  r.Use(loggingMiddleware)

  r.HandleFunc("/identity/login", handler.Login).Methods("POST")
  r.HandleFunc("/identity/register", handler.Register).Methods("POST")

  log.Println("Identity service running on :8081")
  http.ListenAndServe(":8081", r)
}
