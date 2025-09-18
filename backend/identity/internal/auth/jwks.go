package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"` // RSA modulus
	E   string `json:"e"` // RSA exponent
}

// FetchJWKS retrieves the JWKS from Supabase
func FetchJWKS(jwksURL string) (*JWKS, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("JWKS fetch returned status %d", res.StatusCode)
	}

	var set JWKS
	if err := json.NewDecoder(res.Body).Decode(&set); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %w", err)
	}

	if len(set.Keys) == 0 {
		return nil, errors.New("JWKS contains no keys")
	}

	return &set, nil
}

// FindJWKByKeyID searches the JWKS for a key with the matching kid.
func FindJWKByKeyID(jwks *JWKS, kid string) (*JWK, error) {
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			return &key, nil
		}
	}
	return nil, fmt.Errorf("no matching key found for kid: %s", kid)
}

// ConvertJWKToRSAPublicKey takes a JWK and returns an rsa.PublicKey.
func ConvertJWKToRSAPublicKey(jwk JWK) (*rsa.PublicKey, error) {
	// Decode the base64url-encoded modulus (n)
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %w", err)
	}
	n := new(big.Int).SetBytes(nBytes)

	// Decode the base64url-encoded exponent (e)
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %w", err)
	}
	e := new(big.Int).SetBytes(eBytes).Int64()

	// Construct the rsa.PublicKey
	pubKey := &rsa.PublicKey{
		N: n,
		E: int(e),
	}

	return pubKey, nil
}
