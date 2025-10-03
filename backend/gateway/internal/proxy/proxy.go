package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) http.Handler {
	parsedURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("[PROXY] Invalid target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	// Log outbound request
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		log.Printf("[PROXY] → %s %s", req.Method, req.URL.String())
	}

	// Log response status
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf("[PROXY] ← %d %s", resp.StatusCode, resp.Request.URL.String())
		return nil
	}

	// Log errors
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[PROXY] ✖ Error forwarding to %s: %v", target, err)
		http.Error(w, "Gateway error", http.StatusBadGateway)
	}

	return proxy
}
