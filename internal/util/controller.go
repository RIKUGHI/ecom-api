package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/sirupsen/logrus"
)

func HandleApiError(ctx *gin.Context, err error, logMessage string, log *logrus.Logger) {
	apiErr, ok := err.(*ApiError)
	if ok {
		log.Warnf(logMessage, apiErr)
		ctx.JSON(apiErr.Code, model.ApiResponse[*model.UserResponse]{
			Errors: apiErr.Error(),
		})
	} else {
		log.Warnf(logMessage, err)
		ctx.JSON(http.StatusInternalServerError, model.ApiResponse[*model.UserResponse]{
			Errors: err.Error(),
		})
	}
}
