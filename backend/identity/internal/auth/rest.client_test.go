package auth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

type fakePoster struct {
    resp *http.Response
    err  error
}

func (f *fakePoster) Do(req *http.Request) (*http.Response, error) { return f.resp, f.err }

func TestPostJSON_Success(t *testing.T) {
    body := `{"access_token":"abc"}`
    f := &fakePoster{
        resp: &http.Response{
            StatusCode: 200,
            Body:       io.NopCloser(bytes.NewBufferString(body)),
        },
    }
    var out struct{ AccessToken string `json:"access_token"` }
    err := PostJSON(context.Background(), f, "http://example", nil, map[string]string{"k": "v"}, &out)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if out.AccessToken != "abc" {
        t.Fatalf("expected token abc got %q", out.AccessToken)
    }
}
