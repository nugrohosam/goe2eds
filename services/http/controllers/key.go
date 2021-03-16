package controllers

import (
	"net/http"
	helpers "github.com/nugrohosam/goe2eds/helpers"
	resources "github.com/nugrohosam/goe2eds/services/http/resources/v1"

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
			data := resources.KeyItem{
				PrivateKey: privKey,
				PublicKey: pubKey,
				CertUrl: urlLink,
			}

			c.JSON(http.StatusOK, helpers.ResponseOne(data))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseOne(nil))
		}
	}
}
