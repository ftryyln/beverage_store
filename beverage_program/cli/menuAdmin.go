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

// showMenuAdmin menampilkan menu khusus untuk admin
func showMenuAdmin(reader *bufio.Reader, db *sql.DB, user *entity.User) {
	for {
		// Tampilkan pilihan menu admin
		fmt.Println("\n============================")
		fmt.Println("======== Admin Menu ========")
		fmt.Println("============================")
		fmt.Println("1. View All Products")
		fmt.Println("2. Category Menu")
		fmt.Println("3. Add Product")
		fmt.Println("4. Update Product")
		fmt.Println("5. Delete Product")
		fmt.Println("6. User Report")
		fmt.Println("7. Beverage Report")
		fmt.Println("8. Beverage Category Report")
		fmt.Println("0. Logout")
		fmt.Println("============================")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			// Menampilkan semua produk
			products, err := handler.GetAllProducts(db)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			for _, p := range products {
				fmt.Printf("ID: %d | Name: %s | Price: %.2f | Stock: %d | Category Name: %s\n", p.ID, p.ProductName, p.Price, p.Stock, p.CategoryName)
			}

		case "2":
			// Masuk ke sub-menu kategori
			categoryMenu(reader, db)
			continue

		case "3":
			// Menambahkan produk baru
			fmt.Print("Enter product name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter price: ")
			priceStr, _ := reader.ReadString('\n')
			priceStr = strings.TrimSpace(priceStr)

			fmt.Print("Enter stock: ")
			stockStr, _ := reader.ReadString('\n')
			stockStr = strings.TrimSpace(stockStr)
			fmt.Println("\n============================")

			// Menampilkan semua kategori
			categories, err := handler.GetAllCategories(db)
			if err != nil || len(categories) == 0 {
				fmt.Println("❌ No categories available. Please add a category first.")
				continue
			}
			fmt.Println("Available Categories:")
			for _, c := range categories {
				fmt.Printf("%d. %s\n", c.ID, c.Name)
			}
			fmt.Println("============================")

			// Memilih kategori
			fmt.Print("Choose category ID: ")
			catStr, _ := reader.ReadString('\n')
			catStr = strings.TrimSpace(catStr)

			// Konversi input ke integer
			price, _ := strconv.Atoi(priceStr)
			stock, _ := strconv.Atoi(stockStr)
			categoryID, _ := strconv.Atoi(catStr)

			// Menambahkan produk ke database
			err = handler.AddProduct(db, name, price, stock, categoryID, user.ID)
			if err != nil {
				fmt.Println("❌ Failed to add product:", err)
			} else {
				fmt.Println("✅ Product added successfully.")
			}

		case "4":
			// Memperbarui data produk
			fmt.Print("Enter product ID to update: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			fmt.Print("Enter new product name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter new price: ")
			priceStr, _ := reader.ReadString('\n')
			priceStr = strings.TrimSpace(priceStr)
			price, _ := strconv.Atoi(priceStr)

			fmt.Print("Enter new stock: ")
			stockStr, _ := reader.ReadString('\n')
			stockStr = strings.TrimSpace(stockStr)
			stock, _ := strconv.Atoi(stockStr)

			// Update produk
			err := handler.UpdateProduct(db, id, name, price, stock)
			if err != nil {
				fmt.Println("❌ Failed to update product:", err)
			} else {
				fmt.Println("✅ Product updated successfully.")
			}

		case "5":
			// Menghapus produk
			fmt.Print("Enter the product ID to delete: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			productID, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid product ID.")
				return
			}

			err = handler.DeleteProductByID(db, productID)
			if err != nil {
				fmt.Println("Error deleting product:", err)
			} else {
				fmt.Println("✅ Product successfully deleted.")
			}

		case "6":
			// Menampilkan laporan top user
			fmt.Print("How many top users do you want to see? ")
			limitStr, _ := reader.ReadString('\n')
			limitStr = strings.TrimSpace(limitStr)

			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit <= 0 {
				fmt.Println("Invalid input. Please enter a positive number.")
				return
			}

			err = handler.UserReport(db, limit)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "7":
			// Menampilkan laporan produk terbanyak
			fmt.Print("How many beverage report data do you want to see? ")
			limitStr, _ := reader.ReadString('\n')
			limitStr = strings.TrimSpace(limitStr)

			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit <= 0 {
				fmt.Println("Invalid input. Please enter a positive number.")
				return
			}

			err = handler.ItemReport(db, limit)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "8":
			// Menampilkan laporan kategori minuman
			fmt.Print("How many category report data do you want to see? ")
			limitStr, _ := reader.ReadString('\n')
			limitStr = strings.TrimSpace(limitStr)

			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit <= 0 {
				fmt.Println("Invalid input. Please enter a positive number.")
				return
			}

			err = handler.CategoryReport(db, limit)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "0":
			// Logout dari akun admin
			fmt.Println("Logging out...")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}

// categoryMenu menampilkan menu untuk pengelolaan kategori produk
func categoryMenu(reader *bufio.Reader, db *sql.DB) {
	for {
		fmt.Println("\n=====================")
		fmt.Println("=== Category Menu ===")
		fmt.Println("=====================")
		fmt.Println("1. View Categories")
		fmt.Println("2. Add Category")
		fmt.Println("0. Back")
		fmt.Println("=====================")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			// Menampilkan semua kategori
			categories, err := handler.GetAllCategories(db)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			for _, c := range categories {
				fmt.Printf("ID: %d | Name: %s\n", c.ID, c.Name)
			}
		case "2":
			// Menambahkan kategori baru
			fmt.Print("Enter category name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			err := handler.CreateCategory(db, name)
			if err != nil {
				fmt.Println("❌ Failed to add category:", err)
			} else {
				fmt.Println("✅ Category added successfully.")
			}
		case "0":
			// Kembali ke menu admin
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
