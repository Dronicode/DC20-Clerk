package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"dc20clerk/backend/identity/internal/auth"
)

type JWKSProvider struct {
	mu        sync.RWMutex
	jwks      *auth.JWKS
	jwksURL   string
	lastFetch time.Time
}

// NewJWKSProvider initializes and fetches JWKS from the given URL.
func NewJWKSProvider(jwksURL string) *JWKSProvider {
	p := &JWKSProvider{jwksURL: jwksURL}
	if err := p.refresh(); err != nil {
		panic(fmt.Sprintf("failed to fetch JWKS: %v", err))
	}
	return p
}

// Get returns the cached JWKS.
func (p *JWKSProvider) Get() *auth.JWKS {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.jwks
}

// refresh fetches the JWKS from the remote URL.
func (p *JWKSProvider) refresh() error {
	resp, err := http.Get(p.jwksURL)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	var jwks auth.JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("failed to decode JWKS: %w", err)
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.jwks = &jwks
	p.lastFetch = time.Now()
	return nil
}
