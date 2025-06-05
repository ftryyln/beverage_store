package entity

type UserDetail struct {
	ID          int    `db:"user_detail_id"`
	Name        string `db:"name"`
	UserID      int64  `db:"user_id"`
	PhoneNumber int    `db:"phone_number"`
}
