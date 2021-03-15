package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/goe2eds/helpers"
	requests "github.com/nugrohosam/goe2eds/services/http/requests/v1"
	"github.com/nugrohosam/goe2eds/usecases"

	"github.com/gin-gonic/gin"
)

// MessageHandlerCreate is use
func MessageHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message requests.CreateMessage
		c.ShouldBindJSON(&message)

		validate := helpers.NewValidation()
		if err := validate.Struct(message); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		signature, err := usecases.CreateMessage(message.PrivateKey, []byte(message.Message))
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		data := map[string]interface{}{
			"signature": signature,
		}

		c.JSON(http.StatusBadRequest, helpers.ResponseOne(data))
	}
}

// MessageHandlerValidate is use
func MessageHandlerVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message requests.VerifyMessage
		c.ShouldBindJSON(&message)

		validate := helpers.NewValidation()
		if err := validate.Struct(message); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		valid, err := usecases.VerifyMessage(message.PublicKey, message.Signature, []byte(message.Message))
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		data := map[string]interface{}{
			"is_valid": valid,
		}

		c.JSON(http.StatusBadRequest, helpers.ResponseOne(data))
	}
}
