package entity

import "time"

type Product struct {
	ID           int       `db:"product_id"`
	ProductName  string    `db:"product_name"`
	Price        float64   `db:"price"`
	Stock        int       `db:"stock"`
	CategoryName string    `db:"c.name"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    int       `db:"created_by"`
}
