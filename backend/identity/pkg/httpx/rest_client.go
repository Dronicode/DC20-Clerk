package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// HTTPPoster lets callers inject either the real http.Client or a fake in tests.
type HTTPPoster interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient is used when caller passes nil. It has timeouts so requests don't hang forever.
var DefaultHTTPClient HTTPPoster = &http.Client{
	Transport: &http.Transport{
		DialContext:         (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	},
	Timeout: 10 * time.Second,
}

// drainBody reads and closes the body to allow HTTP connection reuse.
func drainBody(r io.ReadCloser) {
	io.Copy(io.Discard, r)
	r.Close()
}

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
	log.Printf("[%s] %s with headers: %v", method, url, headers)

	// Send the request using the injected or default client
	res, err := client.Do(req)
	if err != nil {
		log.Printf("%s request error: %v", method, err)
		return err
	}
	log.Printf("Response status: %d", res.StatusCode)

	// Ensure response body is drained and closed to reuse connections
	defer drainBody(res.Body)

	// If response status is not OK or Created, extract error message from body
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, res.Body)
		if buf.Len() > 0 {
			return errors.New(buf.String())
		}
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
