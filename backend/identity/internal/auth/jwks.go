package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
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
	Kid string `json:"kid"` // Key ID used to match the JWT header
	Kty string `json:"kty"` // Key type: "RSA" or "EC"
	Alg string `json:"alg"` // Algorithm: "RS256", "ES256", etc.
	Use string `json:"use"` // Intended use: usually "sig" for signature

	// RSA fields
	N string `json:"n,omitempty"` // RSA modulus
	E string `json:"e,omitempty"` // RSA exponent

	// EC fields
	Crv string `json:"crv,omitempty"` // Curve name, e.g. "P-256"
	X   string `json:"x,omitempty"`   // X coordinate of EC public key
	Y   string `json:"y,omitempty"`   // Y coordinate of EC public key
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

func ConvertJWKToPublicKey(jwk JWK) (interface{}, error) {
	switch jwk.Kty {
	case "RSA":
		return ConvertJWKToRSAPublicKey(jwk)
	case "EC":
		return ConvertJWKToECPublicKey(jwk)
	default:
		return nil, fmt.Errorf("unsupported key type: %s", jwk.Kty)
	}
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

func ConvertJWKToECPublicKey(jwk JWK) (*ecdsa.PublicKey, error) {
	// Supabase uses P-256 for ES256 tokens
	if jwk.Crv != "P-256" {
		return nil, fmt.Errorf("unsupported curve: %s", jwk.Crv)
	}

	// Decode base64url-encoded X and Y coordinates
	xBytes, err := base64.RawURLEncoding.DecodeString(jwk.X)
	if err != nil {
		return nil, fmt.Errorf("failed to decode X coordinate: %w", err)
	}
	yBytes, err := base64.RawURLEncoding.DecodeString(jwk.Y)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Y coordinate: %w", err)
	}

	// Convert to big.Int for elliptic curve math
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	// Construct the EC public key using P-256 curve
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}
