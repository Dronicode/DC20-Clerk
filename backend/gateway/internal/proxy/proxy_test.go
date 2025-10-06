package proxy

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyToIdentity(t *testing.T) {
	log.Println("[TEST] → Starting TestProxyToIdentity")

	identity := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("identity service reached"))
	}))
	defer identity.Close()

	proxy := NewReverseProxy(identity.URL)

	req := httptest.NewRequest("POST", "/identity/login", nil)
	w := httptest.NewRecorder()

	proxy.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("[TEST] ✖ Expected 200 OK, got %v", res.StatusCode)
	} else {
		log.Println("[TEST] ← Proxy forwarded successfully")
	}
}

func TestProxyServiceUnavailable(t *testing.T) {
	log.Println("[TEST] → Starting TestProxyServiceUnavailable")

	proxy := NewReverseProxy("http://localhost:9999")

	req := httptest.NewRequest("POST", "/identity/login", nil)
	w := httptest.NewRecorder()

	proxy.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadGateway {
		t.Errorf("[TEST] ✖ Expected 502 Bad Gateway, got %v", res.StatusCode)
	} else {
		log.Println("[TEST] ← Proxy correctly handled service unavailability")
	}
}
