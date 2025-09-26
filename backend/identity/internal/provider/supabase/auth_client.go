package supabase

import (
	"context"

	"dc20clerk/backend/identity/pkg/httpx"
)

// TokenResponse models the Supabase token response.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// RegisterUser creates a new user with email/password using Supabase REST /auth/v1/signup.
// Returns nil on success, or an error string from the API on failure.
func RegisterUser(ctx context.Context, client httpx.HTTPPoster, email, password string) error {
	url := authURL("signup")
	req := map[string]string{
		"email":    email,
		"password": password,
	}
	headers := authHeaders("")

	return httpx.DoJSONPost(ctx, client, url, headers, req, nil)
}

// LoginUser requests access and refresh tokens using Supabase REST /auth/v1/token with grant_type=password.
// Returns TokenResponse
func LoginUser(ctx context.Context, client httpx.HTTPPoster, email, password string) (*TokenResponse, error) {
	url := authURL("token?grant_type=password")
	req := map[string]string{
		"email":    email,
		"password": password,
	}
	headers := authHeaders("")

	tokenResp := new(TokenResponse)
	err := httpx.DoJSONPost(ctx, client, url, headers, req, tokenResp)
	return tokenResp, err
}

// LEARN Replaceable function variables to allow tests to inject fakes.
// Production code should call RegisterUserFunc / LoginUserFunc.
var (
	RegisterUserFunc = RegisterUser
	LoginUserFunc    = LoginUser
)
