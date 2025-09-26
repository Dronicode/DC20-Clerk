package auth_test

import (
	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/pkg/util"
	"log"
	"testing"
)

func TestFetchJWKS(t *testing.T) {
	jwksURL := util.Env("SUPABASE_URL") + "auth/v1/.well-known/jwks.json"
	log.Printf("JWKS URL: %s", jwksURL)
	keySet, err := auth.FetchJWKS(jwksURL)
	if err != nil {
		t.Fatalf("Failed to fetch JWKS: %v", err)
	}
	if len(keySet.Keys) == 0 {
		t.Fatal("JWKS returned no keys")
	}
}
