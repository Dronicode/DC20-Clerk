package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyToIdentity(t *testing.T) {
  // Simulate identity service
  identity := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("identity service reached"))
  }))
  defer identity.Close()

  // Create proxy to identity
  proxy := NewReverseProxy(identity.URL)

  // Simulate gateway request
  req := httptest.NewRequest("POST", "/identity/login", nil)
  w := httptest.NewRecorder()

  proxy.ServeHTTP(w, req)

  res := w.Result()
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected 200 OK, got %v", res.StatusCode)
  }
}

func TestProxyServiceUnavailable(t *testing.T) {
  proxy := NewReverseProxy("http://localhost:9999") // no service here

  req := httptest.NewRequest("POST", "/identity/login", nil)
  w := httptest.NewRecorder()

  proxy.ServeHTTP(w, req)

  res := w.Result()
  if res.StatusCode != http.StatusBadGateway {
    t.Errorf("Expected 502 Bad Gateway, got %v", res.StatusCode)
  }
}
