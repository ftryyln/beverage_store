package handler_test

import (
	"beverage_program/config"
	"beverage_program/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	categories, err := handler.GetAllCategories(db)

	assert.NoError(t, err, "Should not error fetching all categories")
	assert.NotNil(t, categories, "Categories slice should not be nil")
	assert.GreaterOrEqual(t, len(categories), 0, "Categories count should be >= 0")

	if len(categories) > 0 {
		// Check fields of the first category for sanity
		c := categories[0]
		assert.Greater(t, c.ID, 0, "Category ID should be greater than 0")
		assert.NotEmpty(t, c.Name, "Category name should not be empty")
	}
}
