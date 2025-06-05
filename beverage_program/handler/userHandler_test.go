package handler_test

import (
	"beverage_program/config"
	"beverage_program/handler"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi hapus user test dari DB agar tidak duplicate saat test ulang
func cleanupTestUser(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM Users WHERE email = ?", email)
	return err
}

func TestLoginUser(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	email := "testloginuser@example.com"
	password := "testpass"
	role := "customer"

	// Hapus user test sebelum insert baru
	err := cleanupTestUser(db, email)
	assert.NoError(t, err, "CleanupTestUser should not return error")

	user, err := handler.RegisterUser(db, email, password, role)
	assert.NoError(t, err, "RegisterUser should not return error")
	assert.NotNil(t, user, "Registered user should not be nil")

	loggedInUser, err := handler.LoginUser(db, email, password)
	assert.NoError(t, err, "LoginUser should not return error with correct credentials")
	assert.NotNil(t, loggedInUser, "Logged in user should not be nil")
	assert.Equal(t, email, loggedInUser.Email)
	assert.Equal(t, role, loggedInUser.Role)

	_, err = handler.LoginUser(db, email, "wrongpassword")
	assert.Error(t, err, "LoginUser should return error with wrong password")
}

func TestRegisterUser(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	email := "newuser@example.com"
	password := "newpass"
	role := "admin"

	// Hapus user test sebelum insert baru
	err := cleanupTestUser(db, email)
	assert.NoError(t, err, "CleanupTestUser should not return error")

	user, err := handler.RegisterUser(db, email, password, role)
	assert.NoError(t, err, "RegisterUser should not return error")
	assert.NotNil(t, user, "Registered user should not be nil")
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
}

func TestRegisterUserDetails(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	email := "detailuser@example.com"
	password := "pass123"
	role := "customer"

	// Hapus user test sebelum insert baru
	err := cleanupTestUser(db, email)
	assert.NoError(t, err, "CleanupTestUser should not return error")

	user, err := handler.RegisterUser(db, email, password, role)
	assert.NoError(t, err, "RegisterUser should not return error")
	assert.NotNil(t, user, "Registered user should not be nil")

	err = handler.RegisterUserDetails(db, user.ID, "John Doe", "08123456789")
	assert.NoError(t, err, "RegisterUserDetails should not return error")
}
