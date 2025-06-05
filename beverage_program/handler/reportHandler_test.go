package handler_test

import (
	"beverage_program/config"
	"beverage_program/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserReport(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	err := handler.UserReport(db, 5)
	assert.NoError(t, err, "UserReport should not return error")
}

func TestItemReport(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	err := handler.ItemReport(db, 5)
	assert.NoError(t, err, "ItemReport should not return error")
}

func TestCategoryReport(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	err := handler.CategoryReport(db, 5)
	assert.NoError(t, err, "CategoryReport should not return error")
}

func TestUserPurchaseHistory(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	// Test dengan user_id yang ada, misal 1
	err := handler.UserPurchaseHistory(db, 1)
	assert.NoError(t, err, "UserPurchaseHistory should not return error")

	// Test dengan user_id yang kemungkinan tidak ada (misal -1)
	err = handler.UserPurchaseHistory(db, -1)
	assert.NoError(t, err, "UserPurchaseHistory with invalid userID should not return error")
}
