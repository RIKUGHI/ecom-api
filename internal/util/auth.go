package util

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret string, userID uint) (string, error) {
	expiration := time.Second * time.Duration(3600*24*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
