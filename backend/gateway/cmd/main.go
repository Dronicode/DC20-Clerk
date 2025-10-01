package main

import (
	"dc20clerk/backend/gateway/internal/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load(".env")

  r := router.NewRouter()

  log.Println("Gateway running on :8080")
  http.ListenAndServe(":8080", r)
}
