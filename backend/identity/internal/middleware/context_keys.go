package middleware

import "net/http"

// contextKey is used to store values in request context safely.
type contextKey string

// UserIDKey is the context key for the authenticated user's ID.
const UserIDKey contextKey = "userID"

// AccessTokenKey is the context key for the authenticated user's access token.
const AccessTokenKey contextKey = "accessToken"

// GetUserID extracts the authenticated user ID from the request context.
func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(UserIDKey).(string)
	return userID, ok
}
