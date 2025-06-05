package handler

import (
	"beverage_program/entity"
	"database/sql"

	// "error"
	"fmt"
	"time"
)

// CreateOrder membuat order baru untuk customer dan menyimpan detail order beserta update stok produk
func CreateOrder(db *sql.DB, customerID int, items []entity.OrderItem) (int, error) {
	// Mulai transaksi
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	// Cek stok untuk setiap item
	for _, item := range items {
		var stock int
		// Ambil stok produk berdasarkan product_id
		err := tx.QueryRow("SELECT stock FROM Products WHERE product_id = ?", item.ProductID).Scan(&stock)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		// Jika stok tidak mencukupi, batalkan transaksi
		if stock < item.Quantity {
			tx.Rollback()
			// return 0, errors.New("Not enough stock for product ID " + string(rune(item.ProductID)))
			return 0, fmt.Errorf("Not enough stock for product ID %d", item.ProductID)
		}
	}

	// Hitung total harga dari semua item
	total := CalculateOrderTotal(items)

	// Masukkan order baru ke tabel Orders
	orderQuery := "INSERT INTO Orders (user_id, order_date, total_amount) VALUES (?, ?, ?)"
	res, err := tx.Exec(orderQuery, customerID, time.Now(), total)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Ambil ID dari order yang baru saja dibuat
	orderID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Query untuk memasukkan ke tabel OrderItems dan update stok produk
	itemQuery := "INSERT INTO OrderItems (order_id, product_id, quantity, subtotal) VALUES (?, ?, ?, ?)"
	updateStockQuery := "UPDATE Products SET stock = stock - ? WHERE product_id = ?"

	// Simpan setiap item dan update stoknya
	for _, item := range items {
		_, err := tx.Exec(itemQuery, orderID, item.ProductID, item.Quantity, item.Subtotal)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		_, err = tx.Exec(updateStockQuery, item.Quantity, item.ProductID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// Commit transaksi jika semua berhasil
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	// Return ID dari order yang baru dibuat
	return int(orderID), nil
}

// CalculateOrderTotal menjumlahkan semua subtotal item dalam satu order
func CalculateOrderTotal(items []entity.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Subtotal
	}
	return total
}
