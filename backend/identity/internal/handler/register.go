package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/service/identity"
)

// Register is the HTTP handler for POST /identity/register
func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Register request: %s %s", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		log.Printf("Registration error: %v", err)
		return
	}
	defer r.Body.Close()

	var req identity.RegisterRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("Registration error: %v", err)
		return
	}

	resp, err := identity.RegisterUser(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Registration error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
