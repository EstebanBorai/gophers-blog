package error_handler

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type APIError struct {
  Status int
  Message string
}

func APIError(status int, message string) *APIError {
  return &APIError {
    Status: status,
    Message: message,
  }
}
