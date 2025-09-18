package handler

import (
	"encoding/json"
	"net/http"

	"dc20clerk/backend/identity/internal/middleware"
)

// ProfileHandler returns basic info about the authenticated user.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from context
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Respond with user ID
	resp := map[string]string{"user_id": userID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
