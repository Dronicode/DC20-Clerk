package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) http.Handler {
  url, _ := url.Parse(target)
  return httputil.NewSingleHostReverseProxy(url)
}
