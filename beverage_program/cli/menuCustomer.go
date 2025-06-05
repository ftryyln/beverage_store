package cli

import (
	"beverage_program/entity"
	"beverage_program/handler"
	"bufio"
	"database/sql"
	"fmt"
)

// showMenuCustomer menampilkan menu utama untuk pengguna dengan role customer
func showMenuCustomer(reader *bufio.Reader, db *sql.DB, user *entity.User) {
	for {
		// Tampilkan pilihan menu customer
		fmt.Println("\n=============================")
		fmt.Println("========= Main Menu =========")
		fmt.Println("=============================")
		fmt.Println("1. View All Products")
		fmt.Println("2. Create Order")
		fmt.Println("3. View My Purchase History")
		fmt.Println("0. Logout")
		fmt.Println("=============================")
		fmt.Print("Choose option: ")

		// Baca input dari user
		var choice string
		fmt.Fscanln(reader, &choice)

		switch choice {
		case "1":
			// Menampilkan semua produk yang tersedia
			products, err := handler.GetAllProducts(db)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			// Tampilkan daftar produk
			for _, p := range products {
				fmt.Printf("ID: %d | Name: %s | Price: %2.f | Stock: %d\n", p.ID, p.ProductName, p.Price, p.Stock)
			}
		case "2":
			// Memulai proses pemesanan produk
			PlaceOrderFlow(reader, db, user)
		case "3":
			// Menampilkan riwayat pembelian user
			err := handler.UserPurchaseHistory(db, user.ID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "0":
			// Logout dari akun
			fmt.Println("Logging out...")
			return
		default:
			// Input tidak valid
			fmt.Println("Invalid option.")
		}
	}
}
