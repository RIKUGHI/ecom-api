package model

import "time"

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=3,max=130"`
}