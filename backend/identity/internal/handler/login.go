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
	log.Printf("[IDENTITY] → %s %s", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[IDENTITY] ✖ Read body: %v", err)
		http.Error(w, "unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req identity.LoginRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("[IDENTITY] ✖ Invalid JSON: %v", err)
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	resp, err := identity.LoginUser(ctx, req)
	if err != nil {
		log.Printf("[IDENTITY] ✖ Login failed: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	log.Printf("[IDENTITY] ← 200 %s", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
