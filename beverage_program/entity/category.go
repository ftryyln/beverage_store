package entity

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
