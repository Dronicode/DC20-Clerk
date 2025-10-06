package handler

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/internal/types"
	"dc20clerk/backend/identity/pkg/httpx"
)

func TestRegister_Success(t *testing.T) {
	log.Println("[TEST] → Starting TestRegister_Success")

	orig := supabase.RegisterUserFunc
	defer func() { supabase.RegisterUserFunc = orig }()

	supabase.RegisterUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) error {
		if email != "a@b.c" || password != "p" {
			t.Fatalf("[TEST] ✖ Unexpected args: %q %q", email, password)
		}
		return nil
	}

	body := `{"email":"a@b.c","password":"p"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Register(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("[TEST] ✖ Expected 201 Created, got %d", res.StatusCode)
	}

	log.Println("[TEST] ← Register success verified")
}

func TestRegister_BadRequest(t *testing.T) {
	log.Println("[TEST] → Starting TestRegister_BadRequest")

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`{}`))
	w := httptest.NewRecorder()

	Register(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("[TEST] ✖ Expected 400 Bad Request, got %d", res.StatusCode)
	}

	log.Println("[TEST] ← Bad request correctly rejected")
}

func TestRegister_SupabaseError(t *testing.T) {
	log.Println("[TEST] → Starting TestRegister_SupabaseError")

	orig := supabase.RegisterUserFunc
	defer func() { supabase.RegisterUserFunc = orig }()

	supabase.RegisterUserFunc = func(ctx context.Context, client httpx.HTTPPoster, email, password string) error {
		if email != "a@b.c" || password != "p" {
			t.Fatalf("[TEST] ✖ Unexpected args: %q %q", email, password)
		}
		return &types.TestErr{S: "upstream failure"}
	}

	body := `{"email":"a@b.c","password":"p"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Register(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("[TEST] ✖ Expected 500 Internal Server Error, got %d", res.StatusCode)
	}

	log.Println("[TEST] ← Supabase error correctly handled")
}
