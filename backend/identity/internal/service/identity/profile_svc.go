package identity

import (
	"context"
	"errors"

	"dc20clerk/backend/identity/internal/provider/supabase"
	"dc20clerk/backend/identity/pkg/httpx"
)

type ProfileResponse struct {
	ID               string `json:"id"`
	Email            string `json:"email"`
	EmailConfirmedAt string `json:"email_confirmed_at"`
	Role             string `json:"role"`
	Aud              string `json:"aud"`
}

// FetchUserProfile retrieves the user's profile using their access token and user ID.
// Returns a structured response or an error.
func FetchUserProfile(ctx context.Context, accessToken string, userID string) (*ProfileResponse, error) {
	if accessToken == "" {
		return nil, errors.New("[PROFILE] missing access token")
	}
	if userID == "" {
		return nil, errors.New("[PROFILE] missing user ID")
	}

	profile, err := supabase.GetUserProfile(ctx, httpx.DefaultHTTPClient, accessToken)
	if err != nil {
		return nil, err
	}

	return &ProfileResponse{
		ID:               profile.ID,
		Email:            profile.Email,
		EmailConfirmedAt: profile.EmailConfirmedAt,
		Role:             profile.Role,
		Aud:              profile.Aud,
	}, nil
}
