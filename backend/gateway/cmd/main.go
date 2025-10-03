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

	log.Println("[GATEWAY] ← Service ready on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("[GATEWAY] ✖ Server failed: %v", err)
	}
}
