package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTestJWT creates a signed JWT and its corresponding JWK.
func GenerateTestJWT() (string, JWK, error) {
	// Step 1: Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", JWK{}, err
	}
	publicKey := &privateKey.PublicKey

	// Step 2: Create JWT claims
	claims := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "Luffy",
		"iat":  time.Now().Unix(),
	}

	// Step 3: Create JWT header with kid
	kid := "test-key"
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid

	// Step 4: Sign the token
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", JWK{}, err
	}

	// Step 5: Convert RSA public key to JWK
	n := base64.RawURLEncoding.EncodeToString(publicKey.N.Bytes())
	e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(publicKey.E)).Bytes())

	jwk := JWK{
		Kid: kid,
		Kty: "RSA",
		Alg: "RS256",
		Use: "sig",
		N:   n,
		E:   e,
	}

	return signedToken, jwk, nil
}
