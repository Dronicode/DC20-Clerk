package handler_test

import (
	"context"
	"dc20clerk/backend/identity/internal/handler"
	"dc20clerk/backend/identity/internal/middleware"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfileHandler_ValidUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/identity/profile", nil)
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, "user-123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", rr.Code)
	}

	var resp map[string]string
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp["user_id"] != "user-123" {
		t.Errorf("Unexpected user_id: %v", resp["user_id"])
	}
}

func TestProfileHandler_MissingUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/identity/profile", nil)
	rr := httptest.NewRecorder()

	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected 401 Unauthorized, got %d", rr.Code)
	}
}

func TestProfileHandler_InvalidUserType(t *testing.T) {
	req := httptest.NewRequest("GET", "/identity/profile", nil)
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, 12345) // not a string
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected 401 Unauthorized, got %d", rr.Code)
	}
}
