package controllers

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/goe2eds/helpers"
	requests "github.com/nugrohosam/goe2eds/services/http/requests/v1"
	"github.com/nugrohosam/goe2eds/usecases"

	resources "github.com/nugrohosam/goe2eds/services/http/resources/v1"
	"github.com/gin-gonic/gin"
)

// FileHandlerCreate is use
func FileHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file requests.CreateFile
		c.ShouldBind(&file)

		validate := helpers.NewValidation()
		if err := validate.Struct(file); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		fileByte, err := helpers.ReadFileRequest(file.File)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		certByte, err := helpers.ReadFileRequest(file.Cert)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		formatedString := fmt.Sprintf("%s", strings.Replace(file.PrivateKey, `\n`, "\n", -1))
		fileUrl, err := usecases.CreateFile(formatedString, fileByte, certByte)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		data := resources.SignatureFileItem{
			FileUrl: fileUrl,
		}

		c.JSON(http.StatusBadRequest, helpers.ResponseOne(data))
	}
}

// FileHandlerValidate is use
func FileHandlerVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file requests.VerifyFile
		c.ShouldBind(&file)

		validate := helpers.NewValidation()
		if err := validate.Struct(file); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		signByte, err := helpers.ReadFileRequest(file.SignatureFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		fileByte, err := helpers.ReadFileRequest(file.File)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
		
		formatedString := fmt.Sprintf("%s", strings.Replace(file.PublicKey, `\n`, "\n", -1))
		valid, err := usecases.VerifyFile(formatedString, signByte, fileByte)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		data := resources.FileItem{
			IsValid: valid,
		}

		c.JSON(http.StatusBadRequest, helpers.ResponseOne(data))
	}
}
