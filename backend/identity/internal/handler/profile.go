package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/middleware"
	"dc20clerk/backend/identity/internal/service/identity"
)

// ProfileHandler returns basic info about the authenticated user.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[PROFILE] ProfileHandler invoked")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Extract user ID from context
	userID, ok := middleware.GetUserID(r)
	if !ok || userID == "" {
		log.Println("[PROFILE] Missing user ID in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	log.Printf("[PROFILE] Extracted user ID: %s\n", userID)

	accessToken, ok := r.Context().Value(middleware.AccessTokenKey).(string)
	if !ok || accessToken == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	resp, err := identity.FetchUserProfile(ctx, accessToken, userID)
	if err != nil {
		log.Printf("[PROFILE] Error fetching profile: %v", err)
		http.Error(w, "failed to fetch profile", http.StatusInternalServerError)
		return
	}

	// Forward response to frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
