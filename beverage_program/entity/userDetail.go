package entity

type UserDetails struct {
	CustomerDetailID int
	CustomerID       int // foreign key ke Users.ID
	Name             string
	PhoneNumber      string
}
