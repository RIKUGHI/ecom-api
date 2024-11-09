package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
)

func NewAuth(log *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			ctx.JSON(http.StatusUnauthorized, model.ApiResponse[*model.UserResponse]{
				Errors: "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		jwtToken := strings.Split(authToken, " ")[1]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			log.Warnf("Failed to get token : %+v", err)
			ctx.JSON(http.StatusUnauthorized, model.ApiResponse[*model.UserResponse]{
				Errors: "Invalid token",
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expirationTime, ok := claims["expiresAt"].(float64)
			if !ok || time.Unix(int64(expirationTime), 0).Before(time.Now()) {
				ctx.JSON(http.StatusUnauthorized, model.ApiResponse[*model.UserResponse]{
					Errors: "Token has expired",
				})
				ctx.Abort()
				return
			}

			if userID, ok := claims[util.UserKey].(string); ok {
				ctx.Set(util.UserKey, userID)
			}
		} else {
			log.Warnf("Failed to get token : %+v", err)
			ctx.JSON(http.StatusUnauthorized, model.ApiResponse[*model.UserResponse]{
				Errors: "Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
