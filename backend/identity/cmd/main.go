package main

import (
	"dc20clerk/backend/identity/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load(".env")

  r := mux.NewRouter()
  r.HandleFunc("/identity/login", handler.Login).Methods("POST")
  r.HandleFunc("/identity/register", handler.Register).Methods("POST")

  log.Println("Identity service running on :8081")
  http.ListenAndServe(":8081", r)
}
