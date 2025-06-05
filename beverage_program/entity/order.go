package entity

import "time"

type Order struct {
	ID          int       `db:"id"`
	CustomerID  int       `db:"customer_id"`
	OrderDate   time.Time `db:"order_date"`
	TotalAmount float64   `db:"total_amount"`
}
