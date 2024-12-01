package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Own error type
type error struct {
	Message string `json:"message"`
}

// newErrorResponse sends an error response to the client.
// Returns:
// - Aborts the request with the given status code and error message
func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, error{Message: message})
}
