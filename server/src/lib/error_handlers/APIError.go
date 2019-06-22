package error_handlers

import (
	"github.com/gin-gonic/gin"
)

// APIError Represents an API Error
type APIError struct {
	Status  int
	Message string
	Stack   string
}

func responseWithError(c *gin.Context, err APIError) {
	c.AbortWithStatusJSON(err.Status, APIError{
		Status:  err.Status,
		Message: err.Message,
		Stack:   err.Stack,
	})
}

func ResponseWithError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, APIError{
		Status:  status,
		Message: message,
	})

	c.Abort()
}

func BadRequest(c *gin.Context, message string) {
	var statusCode int = 400

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})

	c.Abort()
}

func Unauthorized(c *gin.Context, message string) {
	var statusCode int = 401

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})

	c.Abort()
}

func NotFound(c *gin.Context, message string) {
	var statusCode int = 404

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})

	c.Abort()
}

func InternalServerError(c *gin.Context, message string) {
	responseWithError(c, APIError{
		Status:  500,
		Message: message,
	})
}

func SpecialError(c *gin.Context, message string, err error) {
	responseWithError(c, APIError{
		Status:  500,
		Message: message,
		Stack:   err.Error(),
	})
}
