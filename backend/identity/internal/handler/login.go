package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"dc20clerk/backend/identity/internal/supabase"
)

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type loginResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

// Login is the HTTP handler for POST /identity/login
func Login(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "unable to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    var req loginRequest
    if err := json.Unmarshal(body, &req); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Email == "" || req.Password == "" {
        http.Error(w, "email and password required", http.StatusBadRequest)
        return
    }

    tokens, err := supabase.LoginUserFunc(ctx, nil, req.Email, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    resp := loginResponse{
        AccessToken:  tokens.AccessToken,
        RefreshToken: tokens.RefreshToken,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
