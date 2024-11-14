package model

import "time"

type ProductResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

type SearchProductQuery struct {
	Name string `json:"name"`
	Page int    `json:"page" binding:"numeric"`
	Size int    `json:"size" binding:"numeric"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,numeric"`
	Quantity    int     `json:"quantity" binding:"required,numeric"`
}
