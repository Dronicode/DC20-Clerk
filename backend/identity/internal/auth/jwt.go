package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type jwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Kid string `json:"kid"`
}

func ExtractJWTHeader(token string) (*jwtHeader, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT format")
	}

	headerSegment := parts[0]

	// JWT uses base64url encoding (not standard base64)
	decoded, err := base64.RawURLEncoding.DecodeString(headerSegment)
	if err != nil {
		return nil, err
	}

	var header jwtHeader
	if err := json.Unmarshal(decoded, &header); err != nil {
		return nil, err
	}

	return &header, nil
}

// VerifyJWT parses and verifies the JWT using the provided RSA public key.
func VerifyJWT(tokenString string, pubKey *rsa.PublicKey) (*jwt.Token, error) {
	// Parse the token and verify its signature using the public key
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RS256
		if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return pubKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("JWT verification failed: %w", err)
	}

	// Return the verified token
	return token, nil
}

// ValidateToken performs full JWT validation using the JWKS.
func ValidateToken(tokenString string, jwks *JWKS) (*jwt.Token, error) {
	// Step 1: Extract the JWT header to get the kid
	header, err := ExtractJWTHeader(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JWT header: %w", err)
	}

	// Step 2: Find the matching JWK by kid
	key, err := FindJWKByKeyID(jwks, header.Kid)
	if err != nil {
		return nil, fmt.Errorf("failed to find matching JWK: %w", err)
	}

	// Step 3: Convert the JWK to an RSA public key
	pubKey, err := ConvertJWKToRSAPublicKey(*key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert JWK to RSA key: %w", err)
	}

	// Step 4: Verify the JWT signature using the public key
	token, err := VerifyJWT(tokenString, pubKey)
	if err != nil {
		return nil, fmt.Errorf("JWT verification failed: %w", err)
	}

	return token, nil
}
