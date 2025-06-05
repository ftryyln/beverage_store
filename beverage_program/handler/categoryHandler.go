package handler

import (
	"beverage_program/entity"
	"database/sql"
	"fmt"
)

// CreateCategory menambahkan kategori baru ke dalam tabel categories
func CreateCategory(db *sql.DB, name string) error {
	// Query untuk menyisipkan kategori baru
	query := "INSERT INTO Categories (name) VALUES (?)"

	// Eksekusi query dengan parameter nama kategori
	_, err := db.Exec(query, name)
	if err != nil {
		// Jika gagal, kembalikan error dengan pesan
		return fmt.Errorf("❌ Failed to insert category: %w", err)
	}

	// Jika berhasil, kembalikan nil
	return nil
}

// GetAllCategories mengambil semua data kategori dari tabel categories
func GetAllCategories(db *sql.DB) ([]entity.Category, error) {
	// GetAllCategories mengambil semua data kategori dari tabel categories
	query := "SELECT category_id, name FROM Categories"

	// Eksekusi query
	rows, err := db.Query(query)
	if err != nil {
		// Jika gagal mengambil data, kembalikan error
		return nil, fmt.Errorf("❌ Failed to fetch categories: %w", err)
	}
	defer rows.Close()

	// Slice untuk menampung hasil
	var categories []entity.Category

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var c entity.Category

		// Scan data ke struct Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			// Jika gagal membaca data, kembalikan error
			return nil, fmt.Errorf("❌ Failed to scan category: %w", err)
		}

		// Tambahkan ke slice hasil
		categories = append(categories, c)
	}

	// Kembalikan slice hasil
	return categories, nil
}
