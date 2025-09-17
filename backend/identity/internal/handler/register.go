package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/supabase"
)

type registerRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Register is the HTTP handler for POST /identity/register
func Register(w http.ResponseWriter, r *http.Request) {
    // small timeout per request
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    // read body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "unable to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    var req registerRequest
    if err := json.Unmarshal(body, &req); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    // basic validation
    if req.Email == "" || req.Password == "" {
        http.Error(w, "email and password required", http.StatusBadRequest)
        return
    }

    // Call supabase (uses the package-level replaceable function)
    if err := supabase.RegisterUserFunc(ctx, nil, req.Email, req.Password); err != nil {
        // return server error with the message from Supabase (PostJSON returns API body on non-2xx)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // created
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("user created"))
}
