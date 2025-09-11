package handler

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Register endpoint reached"))
}
