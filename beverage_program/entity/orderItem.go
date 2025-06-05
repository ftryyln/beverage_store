package entity

type OrderItem struct {
	ID        int     `db:"id"`
	OrderID   int     `db:"order_id"`
	ProductID int     `db:"product_id"`
	Quantity  int     `db:"quantity"`
	Subtotal  float64 `db:"subtotal"`
}
