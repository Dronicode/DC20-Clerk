package supabase

import (
	"context"
	"dc20clerk/backend/identity/pkg/httpx"
	"log"
)

// UserProfile represents the authenticated user's profile returned by Supabase.
type UserProfile struct {
	ID               string `json:"id"`
	Email            string `json:"email"`
	EmailConfirmedAt string `json:"email_confirmed_at"`
	Role             string `json:"role"`
	Aud              string `json:"aud"`
}

// GetUserProfile fetches the authenticated user's profile from Supabase.
func GetUserProfile(ctx context.Context, client httpx.HTTPPoster, accessToken string) (*UserProfile, error) {
	url := authURL("user")
	headers := authHeaders(accessToken)

	profile := new(UserProfile)
	err := httpx.DoJSONGet(ctx, client, url, headers, profile)
	if err != nil {
		log.Printf("[SUPABASE] ✖ GET %s: %v", url, err)
	}
	return profile, err
}
