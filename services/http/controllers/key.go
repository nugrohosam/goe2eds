package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// KeyHandlerCreate is use
func KeyHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
