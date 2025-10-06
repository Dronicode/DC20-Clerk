package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// doJSONRequest sends an HTTP request with optional JSON body and headers,
// and decodes the JSON response into 'out' (if out != nil).
func doJSONRequest(
	ctx context.Context, // caller's context for timeout/cancel propagation
	client HTTPPoster, // injected HTTP client (or fallback to default)
	method string, // HTTP method: GET, POST, etc.
	url string, // full request URL
	headers map[string]string, // headers like Authorization, apikey
	body io.Reader, // optional request body (nil for GET)
	out interface{}, // optional pointer to decode JSON response into
) error {
	// Use the default client if caller didn't provide one
	if client == nil {
		client = DefaultHTTPClient
	}

	// Construct the HTTP request bound to caller's context
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		log.Printf("[REST] ✖ %s %s: request creation failed: %v", method, url, err)
		return err
	}

	// Set Content-Type for POST requests with JSON body
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	// Apply caller-provided headers (e.g. Authorization, apikey)
	for k, val := range headers {
		req.Header.Set(k, val)
	}

	// Log outbound request
	log.Printf("[REST] → %s %s", method, url)

	// Send the request using the injected or default client
	res, err := client.Do(req)
	if err != nil {
		log.Printf("[REST] ✖ %s %s: request failed: %v", method, url, err)
		return err
	}
	// Ensure response body is drained and closed to reuse connections
	defer drainBody(res.Body)

	log.Printf("[REST] ← %d %s", res.StatusCode, url)

	// If response status is not OK or Created, extract error message from body
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, res.Body)
		if buf.Len() > 0 {
			log.Printf("[REST] ✖ %s %s: error response: %s", method, url, buf.String())
			return errors.New(buf.String())
		}
		log.Printf("[REST] ✖ %s %s: status error: %s", method, url, res.Status)
		return errors.New(res.Status)
	}

	// If caller doesn't want response decoded, return early
	if out == nil {
		return nil
	}

	// Decode JSON response into caller-provided struct
	return json.NewDecoder(res.Body).Decode(out)
}

// DoJSONPost sends a POST request with JSON body and headers.
func DoJSONPost(
	ctx context.Context, client HTTPPoster, url string, headers map[string]string, v interface{}, out interface{},
) error {
	// Marshal request body into JSON bytes
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("[REST] ✖ POST %s: JSON marshal failed: %v", url, err)
		return err
	}

	// Delegate to shared request handler
	return doJSONRequest(ctx, client, http.MethodPost, url, headers, bytes.NewReader(b), out)
}

// DoJSONGet sends a GET request with headers,
func DoJSONGet(
	ctx context.Context, client HTTPPoster, url string, headers map[string]string, out interface{},
) error {
	// Delegate to shared request handler with no body
	return doJSONRequest(ctx, client, http.MethodGet, url, headers, nil, out)
}
