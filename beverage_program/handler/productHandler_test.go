package handler_test

import (
	"beverage_program/config"
	"beverage_program/handler"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {

	// Inisialisasi koneksi ke database (menggunakan konfigurasi dari file .env)
	db := config.InitDB()
	// Pastikan koneksi ditutup setelah test selesai
	defer db.Close()

	// Panggil fungsi GetAllProducts untuk mendapatkan daftar produk
	products, err := handler.GetAllProducts(db)

	// Pastikan tidak terjadi error saat pengambilan data produk
	assert.NoError(t, err, "Should not error when fetching products")

	// Pastikan hasil products tidak bernilai nil
	assert.NotNil(t, products, "Products slice should not be nil")

	// Optional: Pastikan jumlah produk minimal 0 (boleh kosong, tapi tidak error)
	assert.GreaterOrEqual(t, len(products), 0, "Products count should be >= 0")

}

func TestGetProductByID(t *testing.T) {
	db := config.InitDB()
	defer db.Close()

	// You should have at least one known product_id in your test DB
	knownID := 1

	product, err := handler.GetProductByID(db, knownID)

	assert.NoError(t, err, "Should not error fetching product by ID")
	assert.Equal(t, knownID, product.ID, "Product ID should match requested ID")
	assert.NotEmpty(t, product.ProductName, "ProductName should not be empty")
	assert.Greater(t, product.Price, float64(0), "Price should be greater than zero")
	assert.GreaterOrEqual(t, product.Stock, 0, "Stock should be >= 0")
	assert.WithinDuration(t, time.Now(), product.CreatedAt, time.Hour*24*365, "CreatedAt should be within last year")
}
