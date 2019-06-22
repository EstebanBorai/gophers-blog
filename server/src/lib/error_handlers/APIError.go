package error_handlers

import (
	"github.com/gin-gonic/gin"
)

type APIError struct {
	Status  int
	Message string
}

func ResponseWithError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, APIError{
		Status:  status,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	var statusCode int = 400

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})
}

func Unauthorized(c *gin.Context, message string) {
	var statusCode int = 401

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	var statusCode int = 404

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	var statusCode int = 500

	c.AbortWithStatusJSON(statusCode, APIError{
		Status:  statusCode,
		Message: message,
	})
}
