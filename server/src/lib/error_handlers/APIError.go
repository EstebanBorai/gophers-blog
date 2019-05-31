package error_handlers

import (
  "github.com/gin-gonic/gin"
)

type APIError struct {
  Status int
  Message string
}

func ResponseWithError(c *gin.Context, status int, message string) {
  c.AbortWithStatusJSON(status, APIError {
    Status: status,
    Message: message,
  })
}
