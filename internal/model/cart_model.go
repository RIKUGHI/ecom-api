package model

type CartCheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutRequest struct {
	Items []CartCheckoutItem `json:"items" binding:"required"`
}

type CheckoutResponse struct {
	TotalPrice float64 `json:"total_price"`
	OrderID    int     `json:"order_id"`
}
