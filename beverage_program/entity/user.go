package entity

type User struct {
	ID       int    `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
