package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (c *UserController) Login(ctx *gin.Context) {
	request := new(model.LoginUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	response, err := c.UserService.Login(ctx, request)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to login user: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	request := new(model.RegisterUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	response, err := c.UserService.Create(ctx, request)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to register user: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}

func (c *UserController) Current(ctx *gin.Context) {
	authUserID := util.GetUserID(ctx)
	if authUserID == "" {
		c.Log.Warnf("userID is required")
		ctx.JSON(http.StatusUnauthorized, model.ApiResponse[*model.UserResponse]{
			Errors: "Unauthorized",
		})
		return
	}

	request := &model.GetUserRequest{
		ID: authUserID,
	}

	response, err := c.UserService.Current(ctx, request)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to get current user: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}
