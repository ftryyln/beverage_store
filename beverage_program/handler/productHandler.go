package handler

import (
	"beverage_program/entity"
	"database/sql"
	"fmt"
	"time"
)

// GetAllProducts mengambil semua data produk dari tabel Products dan menggabungkan nama kategori (jika ada)
func GetAllProducts(db *sql.DB) ([]entity.Product, error) {
	rows, err := db.Query(`
	SELECT p.product_id, p.product_name, p.price, p.stock, p.created_at,
	IFNULL(c.name, '') AS category_name
	FROM Products p
	LEFT JOIN ProductCategories pc ON p.product_id = pc.product_id
	LEFT JOIN Categories c ON pc.category_id = c.category_id
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product
		var createdAtRaw []byte
		var categoryName sql.NullString

		// Ambil data per baris
		err := rows.Scan(&p.ID, &p.ProductName, &p.Price, &p.Stock, &createdAtRaw, &categoryName)
		if err != nil {
			return nil, err
		}

		// Parsing waktu dari format string ke time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
		if err != nil {
			return nil, fmt.Errorf("❌ Failed to parse created_at: %w", err)
		}
		p.CreatedAt = parsedTime

		// Set nama kategori jika valid
		if categoryName.Valid {
			p.CategoryName = categoryName.String
		} else {
			p.CategoryName = "(❌ No Category)"
		}

		products = append(products, p)
	}
	return products, nil
}

// AddProduct menambahkan produk baru ke tabel Products dan menghubungkannya ke tabel ProductCategories
func AddProduct(db *sql.DB, name string, price int, stock int, categoryID int, createdBy int) error {

	query := "INSERT INTO Products (product_name, price, stock, created_at, created_by) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, name, price, stock, time.Now(), createdBy)

	if err != nil {
		return fmt.Errorf("❌ Failed to insert product: %w", err)
	}

	// Ambil ID dari produk yang baru saja ditambahkan
	productID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("❌ Failed to get inserted product ID: %w", err)
	}

	// Hubungkan produk ke kategori
	linkQuery := "INSERT INTO ProductCategories (product_id, category_id) VALUES (?, ?)"
	_, err = db.Exec(linkQuery, productID, categoryID)
	if err != nil {
		return fmt.Errorf("❌ Failed to link product to category: %w", err)
	}

	return nil
}

// UpdateProduct memperbarui data produk berdasarkan ID
func UpdateProduct(db *sql.DB, id int, name string, price int, stock int) error {
	query := `
		UPDATE Products
		SET product_name = ?, price = ?, stock = ?
		WHERE product_id = ?
	`
	result, err := db.Exec(query, name, price, stock, id)
	if err != nil {
		return fmt.Errorf("❌ Failed to update product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("❌ Failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("❌ No product found with id %d", id)
	}

	return nil
}

// GetProductByID mengambil satu data produk berdasarkan ID-nya
func GetProductByID(db *sql.DB, id int) (entity.Product, error) {
	var p entity.Product
	var createdAtRaw []byte

	query := "SELECT product_id, product_name, price, stock,  created_at FROM Products WHERE product_id = ?"
	row := db.QueryRow(query, id)

	err := row.Scan(&p.ID, &p.ProductName, &p.Price, &p.Stock, &createdAtRaw)
	if err != nil {
		return p, err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
	if err != nil {
		return p, fmt.Errorf("❌ Failed to parse created_at: %w", err)
	}
	p.CreatedAt = parsedTime

	return p, nil
}

func DeleteProductByID(db *sql.DB, productID int) error {
	// Step 1: Hapus dulu relasi di ProductCategories
	_, err := db.Exec("DELETE FROM ProductCategories WHERE product_id = ?", productID)
	if err != nil {
		return fmt.Errorf("❌ Failed to Delete product-Category Relation: %v", err)
	}

	// Step 2: Baru hapus dari Products
	result, err := db.Exec("DELETE FROM Products WHERE product_id = ?", productID)
	if err != nil {
		return fmt.Errorf("❌ Failed to Delete product: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("❌ No Product Found with ID %d", productID)
	}

	return nil
}
