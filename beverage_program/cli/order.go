package cli

import (
	"beverage_program/entity"
	"beverage_program/handler"
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// PlaceOrderFlow menangani alur pemesanan produk oleh customer
func PlaceOrderFlow(reader *bufio.Reader, db *sql.DB, user *entity.User) {
	// Minta input dari user dalam format productID:quantity, pisah dengan koma
	fmt.Println("Input your order as productID:quantity, separated by commas (e.g. 1:2,3:1) :")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Pisahkan input berdasarkan koma menjadi slice berisi pasangan ID:jumlah
	pairs := strings.Split(input, ",")
	var items []entity.OrderItem

	for _, pair := range pairs {
		// Pisahkan setiap pasangan berdasarkan titik dua
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid format:", pair)
			return
		}

		// Ubah string menjadi integer (product ID dan jumlah)
		productID, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		quantity, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil || quantity <= 0 {
			fmt.Println("Invalid product ID or quantity:", pair)
			return
		}

		// Ambil data produk dari database berdasarkan ID
		product, err := handler.GetProductByID(db, productID)
		if err != nil {
			fmt.Println("Product not found:", productID)
			return
		}

		// Hitung subtotal harga (harga * jumlah)
		subtotal := product.Price * float64(quantity)

		// Tambahkan ke dalam daftar item pesanan
		items = append(items, entity.OrderItem{
			ProductID: productID,
			Quantity:  quantity,
			Subtotal:  subtotal,
		})
	}

	// Buat order baru di database dengan data item yang sudah disusun
	orderID, err := handler.CreateOrder(db, user.ID, items)
	if err != nil {
		fmt.Println("❌ Failed to create order:", err)
		return
	}

	// Tampilkan notifikasi sukses
	fmt.Printf("✅ Order successfully created with Order ID: %d\n", orderID)
}
