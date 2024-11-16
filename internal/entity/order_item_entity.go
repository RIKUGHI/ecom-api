package entity

type OrderItem struct {
	ID        int     `gorm:"column:id;primaryKey"`
	OrderID   int     `gorm:"column:orderId"`
	ProductID int     `gorm:"column:productId"`
	Quantity  int     `gorm:"column:quantity"`
	Price     float64 `gorm:"column:price"`
}

func (oi *OrderItem) TableName() string {
	return "order_items"
}
