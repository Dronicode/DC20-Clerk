package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
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

// PostJSON sends v as JSON to url using client (or DefaultHTTPClient when client==nil),
// sets headers, and decodes the JSON response into out (if out != nil).
func PostJSON(ctx context.Context, client HTTPPoster, url string, headers map[string]string, v interface{}, out interface{}) error {
    // Use the default client if caller didn't provide one
    if client == nil {
        client = DefaultHTTPClient
    }

    // Convert the request body (v) to JSON bytes
    b, err := json.Marshal(v)
    if err != nil {
        return err
    }

    // Build an HTTP POST request bound to the caller's context (for cancel/timeouts)
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
    if err != nil {
        return err
    }

    // Required header for JSON body
    req.Header.Set("Content-Type", "application/json")
    // Add any additional headers the caller provided (e.g., apikey, Authorization)
    for k, val := range headers {
        req.Header.Set(k, val)
    }

    // Send the request
    res, err := client.Do(req)
    if err != nil {
        return err
    }
    // Ensure the body is fully read and closed so connections are reused
    defer drainBody(res.Body)

    // If the server returned an error status, try to return the response body as the error text
    if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
        var buf bytes.Buffer
        _, _ = io.Copy(&buf, res.Body)
        if buf.Len() > 0 {
            return errors.New(buf.String())
        }
        return errors.New(res.Status)
    }

    // Caller doesn't want the response body decoded
    if out == nil {
        return nil
    }

    // Decode JSON response into out (out must be a pointer)
    return json.NewDecoder(res.Body).Decode(out)
}

// drainBody reads and closes the body to allow HTTP connection reuse.
func drainBody(r io.ReadCloser) {
    io.Copy(io.Discard, r)
    r.Close()
}
