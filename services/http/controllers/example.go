package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gocashier/helpers"
)

// ContHandlerExample is use
func ContHandlerExample() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(dataResponse))
	}
}
