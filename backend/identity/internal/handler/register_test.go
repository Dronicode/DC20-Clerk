package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
  req := httptest.NewRequest("POST", "/identity/register", nil)
  w := httptest.NewRecorder()

  Register(w, req)

  res := w.Result()
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected status 200, got %v", res.StatusCode)
  }
}