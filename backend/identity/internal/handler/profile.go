package handler

import (
	"io"
	"log"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/middleware"
	"dc20clerk/backend/identity/pkg/utilities"
)

// ProfileHandler returns basic info about the authenticated user.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[PROFILE] ProfileHandler invoked")

	// Extract user ID from context
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	log.Printf("[PROFILE] Extracted user ID: %s\n", userID)
	if !ok || userID == "" {
		log.Println("[PROFILE] Missing user ID in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	log.Println("[PROFILE] Extracted user ID:", userID)

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("[PROFILE] Missing Authorization header")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Call Supabase /auth/v1/user
	log.Println("[PROFILE] Sending GET to " + utilities.Env("SUPABASE_URL") + "auth/v1/user")
	req, err := http.NewRequest("GET", utilities.Env("SUPABASE_URL")+"auth/v1/user", nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("apikey", utilities.Env("SUPABASE_SECRET_KEY"))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[PROFILE] Supabase /user request failed: %v", err)
		http.Error(w, "failed to fetch user profile", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("[PROFILE] Supabase /user returned status: %d\n", resp.StatusCode)
		http.Error(w, "failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	// Forward Supabase response to frontend
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
