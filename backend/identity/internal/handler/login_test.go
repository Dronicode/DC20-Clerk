package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/internal/types"
	"dc20clerk/backend/identity/pkg/httpx"
)

func TestLogin_Success(t *testing.T) {
	log.Println("[TEST] → Starting TestLogin_Success")
	orig := supabase.LoginUserFunc
	defer func() { supabase.LoginUserFunc = orig }()

	supabase.LoginUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) (*supabase.TokenResponse, error) {
		if email != "a@b.c" || password != "p" {
			t.Fatalf("[TEST] ✖ Unexpected args: %q %q", email, password)
		}
		return &supabase.TokenResponse{
			AccessToken:  "abc123",
			RefreshToken: "xyz789",
		}, nil
	}

	body := `{"email":"a@b.c","password":"p"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Login(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("[TEST] ✖ Expected 200, got %d", res.StatusCode)
	}

	var out map[string]string
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatalf("[TEST] ✖ Invalid JSON response: %v", err)
	}
	if out["access_token"] != "abc123" || out["refresh_token"] != "xyz789" {
		t.Fatalf("[TEST] ✖ Unexpected response body: %v", out)
	}

	log.Println("[TEST] ← Login success verified")
}
func TestLogin_BadRequest(t *testing.T) {
	log.Println("[TEST] → Starting TestLogin_BadRequest")

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{}`))
	w := httptest.NewRecorder()

	Login(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("[TEST] ✖ Expected 400, got %d", res.StatusCode)
	}

	log.Println("[TEST] ← Bad request correctly rejected")
}
func TestLogin_SupabaseError(t *testing.T) {
	log.Println("[TEST] → Starting TestLogin_SupabaseError")

	orig := supabase.LoginUserFunc
	defer func() { supabase.LoginUserFunc = orig }()

	supabase.LoginUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) (*supabase.TokenResponse, error) {
		return nil, &types.TestErr{S: "invalid credentials"}
	}

	body := `{"email":"a@b.c","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Login(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("[TEST] ✖ Expected 401, got %d", res.StatusCode)
	}

	log.Println("[TEST] ← Supabase error correctly handled")
}
