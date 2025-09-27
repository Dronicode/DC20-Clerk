package identity

import (
	"context"
	"errors"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/pkg/httpx"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

// RegisterUser validates input and delegates to Supabase
func RegisterUser(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("[REGISTER] email and password required")
	}

	err := supabase.RegisterUserFunc(ctx, httpx.DefaultHTTPClient, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{Message: "user created"}, nil
}
