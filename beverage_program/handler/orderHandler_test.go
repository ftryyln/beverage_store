package handler

import (
	"beverage_program/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateOrderTotal(t *testing.T) {
	// Simulasi data item pesanan yang terdiri dari 3 produk
	items := []entity.OrderItem{
		{ProductID: 1, Quantity: 2, Subtotal: 30000}, // Item pertama subtotal 30.000
		{ProductID: 2, Quantity: 1, Subtotal: 15000}, // Item kedua subtotal 15.000
		{ProductID: 3, Quantity: 3, Subtotal: 45000}, // Item ketiga subtotal 45.000
	}

	// Total yang diharapkan dari penjumlahan semua subtotal
	expected := 90000.0

	// Hitung total menggunakan fungsi yang diuji
	total := CalculateOrderTotal(items)

	// Bandingkan hasil perhitungan dengan nilai yang diharapkan
	assert.Equal(t, expected, total, "Total order harus sesuai dengan jumlah subtotal semua item")
}
