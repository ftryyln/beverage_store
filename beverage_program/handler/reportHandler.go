package handler

import (
	"database/sql"
	"fmt"
)

// UserReport mencetak daftar user dengan jumlah order terbanyak, dibatasi oleh parameter `limit`
func UserReport(db *sql.DB, limit int) error {
	query := `
		SELECT ud.name, COUNT(o.order_id) AS total_orders
		FROM Orders o
		JOIN Users u ON o.user_id = u.user_id
		JOIN UserDetails ud ON u.user_id = ud.user_id
		GROUP BY ud.name
		ORDER BY total_orders DESC
		LIMIT ?
	`

	// Eksekusi query dengan batas limit
	rows, err := db.Query(query, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Top Users by Order Count:")
	i := 1

	// Loop setiap hasil baris
	for rows.Next() {
		var userName string
		var totalOrders int

		// Ambil data user dan total ordernya
		if err := rows.Scan(&userName, &totalOrders); err != nil {
			return err
		}

		// Tampilkan hasil ke terminal
		fmt.Printf("%d. User Name: %s | Total Orders: %d\n", i, userName, totalOrders)
		i++
	}

	if i == 1 {
		fmt.Println("No users found with orders.")
	}

	return nil
}

// ItemReport mencetak daftar produk (beverage) yang paling banyak dipesan
func ItemReport(db *sql.DB, limit int) error {
	query := `
		SELECT p.product_name, SUM(oi.quantity) AS total_ordered
		FROM OrderItems oi
		JOIN Products p ON oi.product_id = p.product_id
		GROUP BY oi.product_id
		ORDER BY total_ordered DESC
		LIMIT ?;
	`

	rows, err := db.Query(query, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Most Ordered Beverages:")
	i := 1
	for rows.Next() {
		var itemName string
		var qty int
		if err := rows.Scan(&itemName, &qty); err != nil {
			return err
		}
		fmt.Printf("%d. Beverage Name: %s | Total Ordered: %d\n", i, itemName, qty)
		i++
	}

	if i == 1 {
		fmt.Println("No item.")
	}

	return nil
}

// CategoryReport mencetak kategori produk yang paling banyak dipesan berdasarkan total kuantitas
func CategoryReport(db *sql.DB, limit int) error {
	query := `
		SELECT c.name AS category_name, SUM(oi.quantity) AS total_quantity
		FROM OrderItems oi
		JOIN Products p ON oi.product_id = p.product_id
		JOIN ProductCategories pc ON p.product_id = pc.product_id
		JOIN Categories c ON pc.category_id = c.category_id
		GROUP BY c.category_id
		ORDER BY total_quantity DESC
		LIMIT ?;
	`

	rows, err := db.Query(query, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Most Ordered Categories:")
	i := 1
	for rows.Next() {
		var category string
		var sold int
		if err := rows.Scan(&category, &sold); err != nil {
			return err
		}
		fmt.Printf("%d. Category: %s | Total Ordered: %d\n", i, category, sold)
		i++
	}

	if i == 1 {
		fmt.Println("❌ No category.")
	}

	return nil
}

// UserPurchaseHistory menampilkan seluruh riwayat pembelian dari user tertentu berdasarkan user_id
func UserPurchaseHistory(db *sql.DB, userID int) error {
	query := `
		SELECT o.order_id, o.order_date, p.product_name, oi.quantity, oi.subtotal
		FROM Orders o
		JOIN OrderItems oi ON o.order_id = oi.order_id
		JOIN Products p ON oi.product_id = p.product_id
		WHERE o.user_id = ?
		ORDER BY o.order_date DESC
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Your Purchase History:")
	i := 1
	for rows.Next() {
		var orderID int
		var orderDate string
		var productName string
		var quantity int
		var subtotal float64
		if err := rows.Scan(&orderID, &orderDate, &productName, &quantity, &subtotal); err != nil {
			return err
		}
		fmt.Printf("%d. Order #%d | Date: %s | Product: %s | Qty: %d | Subtotal: %.2f\n",
			i, orderID, orderDate, productName, quantity, subtotal)
		i++
	}

	if i == 1 {
		fmt.Println("No purchase history found.")
	}

	return nil
}
