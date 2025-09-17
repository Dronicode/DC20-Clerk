package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/internal/supabase"
)

func TestRegister_Success(t *testing.T) {
    // Replace the supabase.RegisterUserFunc with a fake that asserts inputs and returns nil
    orig := supabase.RegisterUserFunc
    defer func() { supabase.RegisterUserFunc = orig }()

    supabase.RegisterUserFunc = func(ctx context.Context, client auth.HTTPPoster, email, password string) error {
    if email != "a@b.c" || password != "p" {
        t.Fatalf("unexpected args: %q %q", email, password)
    }
    return nil
}

    body := `{"email":"a@b.c","password":"p"}`
    req := httptest.NewRequest(http.MethodPost, "/identity/register", bytes.NewBufferString(body))
    w := httptest.NewRecorder()

    Register(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusCreated {
        t.Fatalf("expected status 201 got %d", res.StatusCode)
    }
}

func TestRegister_BadRequest(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/identity/register", bytes.NewBufferString(`{}`))
    w := httptest.NewRecorder()
    Register(w, req)
    if w.Result().StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400 got %d", w.Result().StatusCode)
    }
}

func TestRegister_SupabaseError(t *testing.T) {
    orig := supabase.RegisterUserFunc
    defer func() { supabase.RegisterUserFunc = orig }()

    supabase.RegisterUserFunc = func(ctx context.Context, client auth.HTTPPoster, email, password string) error {
    if email != "a@b.c" || password != "p" {
        t.Fatalf("unexpected args: %q %q", email, password)
    }
    return &testErr{"upstream failure"};
}

    body := `{"email":"a@b.c","password":"p"}`
    req := httptest.NewRequest(http.MethodPost, "/identity/register", bytes.NewBufferString(body))
    w := httptest.NewRecorder()
    Register(w, req)

    if w.Result().StatusCode != http.StatusInternalServerError {
        t.Fatalf("expected 500 got %d", w.Result().StatusCode)
    }
}

// tiny test error type so we control Error() string
type testErr struct{ s string }
func (e *testErr) Error() string { return e.s }
