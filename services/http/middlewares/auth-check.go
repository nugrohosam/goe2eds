package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/goe2eds/helpers"
	requests "github.com/nugrohosam/goe2eds/services/http/requests/v1"
	requests "github.com/nugrohosam/goe2eds/services/infrastructure"
)

// AuthJwt using for ..
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header requests.HeaderJwt
		c.BindHeader(&header)

		validate := helpers.NewValidation()
		if err := validate.Struct(header); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusUnauthorized, helpers.ResponseErrValidation(fieldsErrors))
			c.Abort()
			return
		}

		token := strings.Replace(header.Authorization, "Bearer ", "", len(header.Authorization))
		if err := usecases.AuthorizationValidation(token); err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr(err.Error()))
			c.Abort()
			return
		}

		userData, _ := usecases.GetDataAuth(token)
		helpers.SetAuth(&helpers.Auth{
			ID:       int(userData["id"].(float64)),
			Name:     userData["name"].(string),
			Username: userData["username"].(string),
			Email:    userData["email"].(string),
		})

		c.Next()
	}
}
