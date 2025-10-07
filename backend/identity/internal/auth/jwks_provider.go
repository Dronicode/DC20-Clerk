package auth

import (
	"dc20clerk/common/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// JWKSProvider caches and refreshes JWKS from a remote URL.
type JWKSProvider struct {
	mu        sync.RWMutex
	jwks      *JWKS
	jwksURL   string
	lastFetch time.Time
}

// NewJWKSProvider initializes the provider and fetches JWKS from the given URL.
func NewJWKSProvider() *JWKSProvider {
	cfg := config.LoadIdentityConfig()
	jwksURL := cfg.SupabaseURL + "auth/v1/.well-known/jwks.json"
	log.Printf("[IDENTITY] → Initializing JWKS provider: %s", jwksURL)
	p := &JWKSProvider{jwksURL: jwksURL}
	if err := p.refresh(); err != nil {
		log.Fatalf("[JWKS] ✖ Initial fetch failed: %v", err)
	}
	return p
}

// Get returns the currently cached JWKS.
func (p *JWKSProvider) Get() *JWKS {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.jwks
}

// refresh fetches and updates the JWKS from the remote URL.
func (p *JWKSProvider) refresh() error {
	log.Printf("[JWKS] → Fetching JWKS from %s", p.jwksURL)

	resp, err := http.Get(p.jwksURL)
	if err != nil {
		return fmt.Errorf("[JWKS] ✖ HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[JWKS] ✖ Unexpected status code: %d", resp.StatusCode)
	}

	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("[JWKS] ✖ JSON decode failed: %w", err)
	}

	if len(jwks.Keys) == 0 {
		return fmt.Errorf("[JWKS] ✖ No keys found in JWKS")
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.jwks = &jwks
	p.lastFetch = time.Now()
	return nil
}

// StartAutoRefresh launches a background goroutine that periodically refreshes the JWKS.
// Recommended interval: 30 minutes or less.
func (p *JWKSProvider) StartAutoRefresh(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			if err := p.refresh(); err != nil {
				log.Printf("[JWKS] ✖ Auto-refresh failed: %v", err)
			} else {
				log.Printf("[JWKS] ← JWKS auto-refreshed")
			}
		}
	}()
}
