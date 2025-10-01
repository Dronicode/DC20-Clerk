package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/pkg/httpx"
)

func TestLogin_Success(t *testing.T) {
	orig := supabase.LoginUserFunc
	defer func() { supabase.LoginUserFunc = orig }()

	supabase.LoginUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) (*supabase.TokenResponse, error) {
		if email != "a@b.c" || password != "p" {
			t.Fatalf("unexpected args: %q %q", email, password)
		}
		return &supabase.TokenResponse{
			AccessToken:  "abc123",
			RefreshToken: "xyz789",
		}, nil
	}

	body := `{"email":"a@b.c","password":"p"}`
	req := httptest.NewRequest(http.MethodPost, "/identity/login", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Login(w, req)

	res := w.Result()
	defer res.Body.Close()
	log.Printf("Response status: %d", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

	b, _ := io.ReadAll(res.Body)
	log.Printf("Raw response body: %s", string(b))
	var out map[string]string
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if out["access_token"] != "abc123" || out["refresh_token"] != "xyz789" {
		t.Fatalf("unexpected response body: %v", out)
	}
}
func TestLogin_BadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/identity/login", bytes.NewBufferString(`{}`))
	w := httptest.NewRecorder()

	Login(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d", res.StatusCode)
	}
}
func TestLogin_SupabaseError(t *testing.T) {
	orig := supabase.LoginUserFunc
	defer func() { supabase.LoginUserFunc = orig }()

	supabase.LoginUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) (*supabase.TokenResponse, error) {
		return nil, &testErr{"invalid credentials"}
	}

	body := `{"email":"a@b.c","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/identity/login", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Login(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 got %d", res.StatusCode)
	}

	b, _ := io.ReadAll(res.Body)
	t.Logf("response body: %s", b)
}
