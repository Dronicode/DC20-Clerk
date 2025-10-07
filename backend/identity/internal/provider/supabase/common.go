package supabase

import (
	"dc20clerk/common/config"
)

var cfg = config.LoadIdentityConfig()

// authURL constructs a full Supabase auth endpoint URL.
//
// Input:
//   - path: the endpoint path segment (e.g. "signup", "token?grant_type=password")
//
// Returns:
//   - full URL string to the Supabase auth endpoint
//
// Example:
//
//	authURL("signup") → "https://xyz.supabase.co/auth/v1/signup"
func authURL(path string) string {
	return cfg.SupabaseURL + "auth/v1/" + path
}

// authHeaders returns the required headers for Supabase auth requests.
//
// Input:
//   - accessToken: the user's access token if available; if empty, falls back to Supabase secret key
//
// Returns:
//   - map of headers including "Authorization" and "apikey"
//
// Behavior:
//   - If accessToken is provided → uses it for Authorization
//   - If accessToken is empty → uses Supabase secret key for Authorization
//
// Example:
//
//	authHeaders("user-token") → Authorization: Bearer user-token
//	authHeaders("")           → Authorization: Bearer <SUPABASE_SECRET_KEY>
func authHeaders(accessToken string) map[string]string {
	token := cfg.SupabaseSecretKey
	if accessToken != "" {
		token = accessToken
	}
	return map[string]string{
		"apikey":        cfg.SupabaseSecretKey,
		"Authorization": "Bearer " + token,
	}
}
