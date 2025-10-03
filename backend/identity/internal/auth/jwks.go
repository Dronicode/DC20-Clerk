package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

// JWKS represents a JSON Web Key Set, typically fetched from an identity provider.
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// JWK represents a single JSON Web Key used to verify JWTs.
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

// FetchJWKS retrieves the JWKS from the given URL and decodes it.
func FetchJWKS(jwksURL string) (*JWKS, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ FetchJWKS: HTTP GET failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[JWKS] ✖ FetchJWKS: unexpected status %d", res.StatusCode)
	}

	var set JWKS
	if err := json.NewDecoder(res.Body).Decode(&set); err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ FetchJWKS: JSON decode failed: %w", err)
	}

	if len(set.Keys) == 0 {
		return nil, fmt.Errorf("[JWKS] ✖ FetchJWKS: no keys found")
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
	return nil, fmt.Errorf("[JWKS] ✖ FindJWKByKeyID: no match for kid=%s", kid)
}

// ConvertJWKToPublicKey converts a JWK to a usable RSA or EC public key.
func ConvertJWKToPublicKey(jwk JWK) (interface{}, error) {
	switch jwk.Kty {
	case "RSA":
		return ConvertJWKToRSAPublicKey(jwk)
	case "EC":
		return ConvertJWKToECPublicKey(jwk)
	default:
		return nil, fmt.Errorf("[JWKS] ✖ ConvertJWKToPublicKey: unsupported key type %s", jwk.Kty)
	}
}

func ConvertJWKToRSAPublicKey(jwk JWK) (*rsa.PublicKey, error) {
	// Decode the base64url-encoded modulus (n)
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ RSA decode modulus: %w", err)
	}

	// Decode the base64url-encoded exponent (e)
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ RSA decode exponent: %w", err)
	}
	e := new(big.Int).SetBytes(eBytes).Int64()

	// Construct and return the rsa.PublicKey
	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(e),
	}, nil
}

func ConvertJWKToECPublicKey(jwk JWK) (*ecdsa.PublicKey, error) {
	// Supabase uses P-256 for ES256 tokens
	if jwk.Crv != "P-256" {
		return nil, fmt.Errorf("[JWKS] ✖ EC unsupported curve: %s", jwk.Crv)
	}

	// Decode base64url-encoded X and Y coordinates
	xBytes, err := base64.RawURLEncoding.DecodeString(jwk.X)
	if err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ EC decode X: %w", err)
	}
	yBytes, err := base64.RawURLEncoding.DecodeString(jwk.Y)
	if err != nil {
		return nil, fmt.Errorf("[JWKS] ✖ EC decode Y: %w", err)
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
