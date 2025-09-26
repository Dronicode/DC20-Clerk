package auth

import (
	"encoding/json"
	"fmt"
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
func NewJWKSProvider(jwksURL string) *JWKSProvider {
	p := &JWKSProvider{jwksURL: jwksURL}
	if err := p.refresh(); err != nil {
		panic(fmt.Sprintf("failed to fetch JWKS: %v", err))
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
	resp, err := http.Get(p.jwksURL)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("failed to decode JWKS: %w", err)
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
				fmt.Printf("[JWKS] Auto-refresh failed: %v\n", err)
			} else {
				fmt.Println("[JWKS] JWKS refreshed successfully")
			}
		}
	}()
}
