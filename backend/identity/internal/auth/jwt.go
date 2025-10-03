package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// jwtHeader represents the decoded JWT header segment.
type jwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Kid string `json:"kid"`
}

// ExtractJWTHeader decodes the JWT header segment to extract algorithm and key ID.
func ExtractJWTHeader(token string) (*jwtHeader, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("[AUTH] ✖ ExtractJWTHeader: invalid JWT format")
	}

	headerSegment := parts[0]

	// JWT uses base64url encoding (not standard base64)
	decoded, err := base64.RawURLEncoding.DecodeString(headerSegment)
	if err != nil {
		return nil, fmt.Errorf("[AUTH] ✖ ExtractJWTHeader: base64 decode failed: %w", err)
	}

	var header jwtHeader
	if err := json.Unmarshal(decoded, &header); err != nil {
		return nil, fmt.Errorf("[AUTH] ✖ ExtractJWTHeader: JSON unmarshal failed: %w", err)
	}

	return &header, nil
}

// VerifyJWT parses and verifies the JWT using the provided public key and algorithm.
func VerifyJWT(tokenString string, pubKey interface{}, alg string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != alg {
			return nil, fmt.Errorf("[AUTH] ✖ VerifyJWT: unexpected signing method: %v", t.Header["alg"])
		}
		return pubKey, nil
	})
}

// ValidateToken performs full JWT validation using the JWKS.
// ValidateToken extracts the header, finds the matching key, converts it, and verifies the signature.
func ValidateToken(tokenString string, jwks *JWKS) (*jwt.Token, error) {
	log.Printf("[AUTH] → Validating token")
	// Step 1: Extract the JWT header to get the kid
	header, err := ExtractJWTHeader(tokenString)
	if err != nil {
		log.Printf("[AUTH] ✖ ExtractJWTHeader failed: %v", err)
		return nil, fmt.Errorf("failed to extract JWT header: %w", err)
	}
	log.Printf("[AUTH] Header: alg=%s kid=%s", header.Alg, header.Kid)

	// Step 2: Find the matching JWK by kid
	key, err := FindJWKByKeyID(jwks, header.Kid)
	if err != nil {
		log.Printf("[AUTH] ✖ FindJWKByKeyID failed: %v", err)
		return nil, fmt.Errorf("failed to find matching JWK: %w", err)
	}

	// Step 3: Convert the JWK to an RSA public key
	pubKey, err := ConvertJWKToPublicKey(*key)
	if err != nil {
		log.Printf("[AUTH] ✖ ConvertJWKToPublicKey failed: %v", err)
		return nil, fmt.Errorf("failed to convert JWK to public key: %w", err)
	}

	// Step 4: Verify the JWT signature using the public key
	token, err := VerifyJWT(tokenString, pubKey, header.Alg)
	if err != nil {
		log.Printf("[AUTH] ✖ VerifyJWT failed: %v", err)
		return nil, fmt.Errorf("JWT verification failed: %w", err)
	}

	log.Printf("[AUTH] ← Token verified successfully")
	return token, nil
}
