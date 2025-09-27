// Centralizes shared HTTP client configuration, interfaces, and connection hygiene.
package httpx

import (
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

// drainBody reads and closes the body to allow HTTP connection reuse.
func drainBody(r io.ReadCloser) {
	io.Copy(io.Discard, r)
	r.Close()
}
