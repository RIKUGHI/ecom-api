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
)
