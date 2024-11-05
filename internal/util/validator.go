package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rikughi/ecom-api/internal/model"
)

func HandleValidationErrors(ctx *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		validationErrors := make(map[string]string)

		for _, fieldError := range errs {
			key := strings.ToLower(string(fieldError.Field()[0])) + fieldError.Field()[1:]
			validationErrors[key] = fmt.Sprintf("Field %s: %s", fieldError.Field(), fieldError.Tag())
		}

		ctx.JSON(http.StatusBadRequest, model.ApiResponse[*model.UserResponse]{
			Errors: validationErrors,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, model.ApiResponse[*model.UserResponse]{
			Errors: err.Error(),
		})
	}
}
