package auth_test

import (
	"dc20clerk/backend/identity/internal/auth"
	"dc20clerk/backend/identity/pkg/util"
	"log"
	"testing"
)

func TestFetchJWKS(t *testing.T) {
	jwksURL := util.Env("SUPABASE_URL") + "auth/v1/.well-known/jwks.json"
	log.Printf("[TEST] → Fetching JWKS from %s", jwksURL)

	keySet, err := auth.FetchJWKS(jwksURL)
	if err != nil {
		t.Fatalf("[TEST] ✖ FetchJWKS failed: %v", err)
	}
	if len(keySet.Keys) == 0 {
		t.Fatal("[TEST] ✖ JWKS returned no keys")
	}

	log.Printf("[TEST] ← JWKS fetch successful: %d keys", len(keySet.Keys))
}
