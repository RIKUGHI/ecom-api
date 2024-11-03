package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	UserService *service.UserService
}

func NewUserController(logger *logrus.Logger, userService *service.UserService) *UserController {
	return &UserController{
		Log:         logger,
		UserService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	request := new(model.RegisterUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
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

		return
	}

	response, err := c.UserService.Create(ctx, request)
	if err != nil {
		apiErr, ok := err.(*util.ApiError)
		if ok {
			c.Log.Warnf("Failed to register user: %+v", apiErr)
			ctx.JSON(apiErr.Code, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: apiErr.Error(),
			})
		} else {
			c.Log.Warnf("Failed to register user: %+v", err)
			ctx.JSON(http.StatusInternalServerError, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}
