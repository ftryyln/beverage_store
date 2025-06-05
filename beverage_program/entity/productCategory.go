package entity

type ProductCategory struct {
	ProductID  int `db:"product_id"`
	CategoryID int `db:"category_id"`
}
