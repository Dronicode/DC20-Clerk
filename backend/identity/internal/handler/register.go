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

	var req identity.RegisterRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("[IDENTITY] ✖ Invalid JSON: %v", err)
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	resp, err := identity.RegisterUser(ctx, req)
	if err != nil {
		log.Printf("[IDENTITY] ✖ Registration failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("[IDENTITY] ← 201 %s", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
