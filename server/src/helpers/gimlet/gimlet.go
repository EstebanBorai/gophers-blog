package gimlet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError represents an error response
type APIError struct {
	StatusCode int
	Message    string
}

// ResponseError represents the error to be sent to the client
type ResponseError struct {
	StatusCode int
	Status     string
	Message    string
}

func end(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, ResponseError{
		StatusCode: status,
		Message:    message,
		Status:     http.StatusText(status),
	})
}

// BadRequest (400) HTTP Response
func BadRequest(c *gin.Context, message string) {
	end(c, http.StatusBadRequest, message)
}

// NotFound (404) HTTP Response
func NotFound(c *gin.Context, message string) {
	end(c, http.StatusNotFound, message)
}

// InternalServerError (500) HttpResponse
func InternalServerError(c *gin.Context, message string) {
	end(c, http.StatusInternalServerError, message)
}
