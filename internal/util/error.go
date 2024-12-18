package util

import (
	"errors"
	"net/http"
)

type ApiError struct {
	Err  error
	Code int
}

func (a *ApiError) Error() string {
	return a.Err.Error()
}

var (
	ErrInternalServer = &ApiError{Err: errors.New("Internal Server Error"), Code: http.StatusInternalServerError}
	ErrUserExists     = &ApiError{Err: errors.New("User already exists"), Code: http.StatusConflict}
	ErrPasswordHash   = &ApiError{Err: errors.New("Failed to generate bcrypt hash"), Code: http.StatusInternalServerError}
	ErrCreateUser     = &ApiError{Err: errors.New("Failed to create user in database"), Code: http.StatusInternalServerError}
	ErrUserNotFound   = &ApiError{Err: errors.New("User not found"), Code: http.StatusNotFound}
	ErrUnauthorized   = &ApiError{Err: errors.New("Unauthorized"), Code: http.StatusUnauthorized}
	ErrInvalidCreds   = &ApiError{Err: errors.New("Invalid credentials"), Code: http.StatusUnauthorized}

	ErrCreateProduct   = &ApiError{Err: errors.New("Failed to create product in database"), Code: http.StatusInternalServerError}
	ErrProductNotFound = &ApiError{Err: errors.New("Product not found"), Code: http.StatusNotFound}

	ErrCreateOrder     = &ApiError{Err: errors.New("Failed to create order in database"), Code: http.StatusInternalServerError}
	ErrUpdateOrder     = &ApiError{Err: errors.New("Failed to update order in database"), Code: http.StatusInternalServerError}
	ErrCreateOrderItem = &ApiError{Err: errors.New("Failed to create order item in database"), Code: http.StatusInternalServerError}
)
