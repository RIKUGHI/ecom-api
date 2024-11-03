package model

type ApiResponse[T any] struct {
	Data   T   `json:"data"`
	Errors any `json:"errors,omitempty"`
}
