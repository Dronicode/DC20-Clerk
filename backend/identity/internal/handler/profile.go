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
	log.Printf("[IDENTITY] → %s %s", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Extract user ID from context
	userID, ok := middleware.GetUserID(r)
	if !ok || userID == "" {
		log.Printf("[IDENTITY] ✖ Missing user ID in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	log.Printf("[PROFILE] Extracted user ID: %s\n", userID)

	accessToken, ok := r.Context().Value(middleware.AccessTokenKey).(string)
	if !ok || accessToken == "" {
		log.Printf("[IDENTITY] ✖ Missing access token in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	resp, err := identity.FetchUserProfile(ctx, accessToken, userID)
	if err != nil {
		log.Printf("[IDENTITY] ✖ Profile fetch failed: %v", err)
		http.Error(w, "failed to fetch profile", http.StatusInternalServerError)
		return
	}

	// Forward response to frontend
	log.Printf("[IDENTITY] ← 200 %s", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
