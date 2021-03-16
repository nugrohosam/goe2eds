package controllers

import (
	"net/http"
	helpers "github.com/nugrohosam/goe2eds/helpers"

	"github.com/nugrohosam/goe2eds/usecases"
	"github.com/gin-gonic/gin"
)

// KeyHandlerCreate is use
func KeyHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		privKey, pubKey, urlLink, err := usecases.CreateKey()

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}
		
		if len(privKey) > 0 && len(pubKey) > 0 {
			data := gin.H{
				"private_key": privKey,
				"public_key": pubKey,
				"url_cert": urlLink,
			}

			c.JSON(http.StatusOK, helpers.ResponseOne(data))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseOne(nil))
		}
	}
}
