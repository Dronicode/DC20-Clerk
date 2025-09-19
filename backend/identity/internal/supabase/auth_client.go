package supabase

import (
	"context"

	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/pkg/utilities"
)

// TokenResponse models the parts of the Supabase token response we care about.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// RegisterUser creates a new user with email/password using Supabase REST /auth/v1/signup.
// Returns nil on success, or an error string from the API on failure.
func RegisterUser(ctx context.Context, client auth.HTTPPoster, email, password string) error {
	url := utilities.Env("SUPABASE_URL") + "auth/v1/signup"
	req := map[string]string{
		"email":    email,
		"password": password,
	}
	headers := map[string]string{
		"apikey":        utilities.Env("SUPABASE_SECRET_KEY"),
		"Authorization": "Bearer " + utilities.Env("SUPABASE_SECRET_KEY"),
	}
	// We don't expect a useful body on success; PostJSON returns decoded body only if out != nil
	return auth.PostJSON(ctx, client, url, headers, req, nil)
}

// LoginUser requests access and refresh tokens using Supabase REST /auth/v1/token with grant_type=password.
func LoginUser(ctx context.Context, client auth.HTTPPoster, email, password string) (*TokenResponse, error) {
	url := utilities.Env("SUPABASE_URL") + "auth/v1/token?grant_type='password'"

	req := map[string]string{
		"email":    email,
		"password": password,
	}
	headers := map[string]string{
		"apikey":        utilities.Env("SUPABASE_SECRET_KEY"),
		"Authorization": "Bearer " + utilities.Env("SUPABASE_SECRET_KEY"),
	}
	var tr TokenResponse
	if err := auth.PostJSON(ctx, client, url, headers, req, &tr); err != nil {
		return nil, err
	}
	return &tr, nil
}

// Replaceable function variables to allow tests to inject fakes.
// Production code should call RegisterUserFunc / LoginUserFunc.
var (
	RegisterUserFunc = RegisterUser
	LoginUserFunc    = LoginUser
)
