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

// Login is the HTTP handler for POST /identity/login
func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login request: %s %s", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		log.Printf("Login error: %v", err)
		return
	}
	defer r.Body.Close()

	var req identity.LoginRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("Login error: %v", err)
		return
	}

	resp, err := identity.LoginUser(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Printf("Login error: %v", err)
		return
	}

	log.Printf("loginResponse: %s\n", resp.AccessToken)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
