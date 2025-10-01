package identity

import (
	"context"
	"errors"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/pkg/httpx"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginUser(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("[LOGIN] email and password required")
	}

	token, err := supabase.LoginUserFunc(ctx, httpx.DefaultHTTPClient, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil

}
