package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
  req := httptest.NewRequest("POST", "/identity/login", nil)
  w := httptest.NewRecorder()

  Login(w, req)

  res := w.Result()
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected status 200, got %v", res.StatusCode)
  }
}