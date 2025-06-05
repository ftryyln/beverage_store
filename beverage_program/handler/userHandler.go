package handler

import (
	"beverage_program/entity"
	"database/sql"
	"fmt"
)

// LoginUser memeriksa email dan password, lalu mengembalikan data user jika cocok
func LoginUser(db *sql.DB, email, password string) (*entity.User, error) {
	var user entity.User

	// Query SQL untuk mencari user berdasarkan email dan password
	query := "SELECT user_id, email, password, role FROM Users WHERE email = ? AND password = ?"

	// Jalankan query dan scan hasil ke struct user
	err := db.QueryRow(query, email, password).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		// Jika tidak ditemukan (email atau password salah), kembalikan error yang jelas
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Invalid email or password")
		}

		// Error selain itu
		return nil, err
	}

	// Login sukses, kembalikan pointer ke user
	return &user, nil
}

// RegisterUser menambahkan user baru ke tabel Users
func RegisterUser(db *sql.DB, email, password, role string) (*entity.User, error) {
	var user entity.User

	// Query untuk menyisipkan user baru
	query := "INSERT INTO Users (email, password, role) VALUES (?, ?, ?)"
	result, err := db.Exec(query, email, password, role)
	if err != nil {
		// Gagal insert
		return nil, err
	}

	// Ambil ID user yang baru dibuat
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set nilai-nilai user yang berhasil didaftarkan
	user.ID = int(lastInsertID)
	user.Email = email
	user.Role = role

	return &user, nil
}

// RegisterUserDetails menambahkan detail user ke tabel UserDetails
func RegisterUserDetails(db *sql.DB, userID int, name, phoneNumber string) error {
	// Query untuk menyimpan nama dan nomor HP berdasarkan user_id
	query := "INSERT INTO UserDetails (user_id, name, phone_number) VALUES  (?, ?, ?)"
	_, err := db.Exec(query, userID, name, phoneNumber)
	// return nil jika sukses, atau error jika gagal
	return err
}
