package handler_test

import (
	"context"
	"dc20clerk/backend/identity/internal/handler"
	"dc20clerk/backend/identity/internal/middleware"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfileHandler_ValidUser(t *testing.T) {
	log.Println("[TEST] → Starting TestProfileHandler_ValidUser")

	req := httptest.NewRequest("GET", "/profile", nil)
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, "user-123")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("[TEST] ✖ Expected 200 OK, got %d", rr.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("[TEST] ✖ Failed to decode response: %v", err)
	}
	if resp["user_id"] != "user-123" {
		t.Errorf("[TEST] ✖ Unexpected user_id: %v", resp["user_id"])
	} else {
		log.Println("[TEST] ← Profile returned correct user_id")
	}
}

func TestProfileHandler_MissingUser(t *testing.T) {
	log.Println("[TEST] → Starting TestProfileHandler_MissingUser")

	req := httptest.NewRequest("GET", "/profile", nil)
	rr := httptest.NewRecorder()

	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("[TEST] ✖ Expected 401 Unauthorized, got %d", rr.Code)
	}

	log.Println("[TEST] ← Missing user correctly rejected")
}

func TestProfileHandler_InvalidUserType(t *testing.T) {
	log.Println("[TEST] → Starting TestProfileHandler_InvalidUserType")

	req := httptest.NewRequest("GET", "/profile", nil)
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, 12345)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler.ProfileHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("[TEST] ✖ Expected 401 Unauthorized, got %d", rr.Code)
	}

	log.Println("[TEST] ← Invalid user type correctly rejected")
}
