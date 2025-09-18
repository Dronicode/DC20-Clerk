package auth_test

import (
	"dc20clerk/backend/identity/internal/auth"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

// This test verifies that ValidateToken correctly parses and verifies a JWT.
func TestValidateToken_Success(t *testing.T) {
	// Step 1: Create a mock JWKS with a known key
	tokenString, jwk, err := auth.GenerateTestJWT()
	if err != nil {
		t.Fatalf("Failed to generate test JWT: %v", err)
	}

	// Step 2: Create jwks to use for token validation
	jwks := &auth.JWKS{Keys: []auth.JWK{jwk}}

	// Step 3: Validate the token
	token, err := auth.ValidateToken(tokenString, jwks)
	if err != nil {
		t.Fatalf("Token validation failed: %v", err)
	}

	// Step 4: Assert claims are present
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("Expected MapClaims, got %T", token.Claims)
	}

	if claims["name"] != "Luffy" {
		t.Errorf("Unexpected claim value: %v", claims["name"])
	}
}
