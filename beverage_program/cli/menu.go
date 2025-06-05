package cli

import (
	"beverage_program/entity"
	"beverage_program/handler"
	"bufio"
	"database/sql"
	"fmt"
	"strings"
)

// ShowMenu menampilkan menu utama untuk login, register, atau keluar
func ShowMenu(reader *bufio.Reader, db *sql.DB) {
	for {
		// Menampilkan menu utama
		fmt.Println("\n==========================")
		fmt.Println("=== Beverage Store CLI ===")
		fmt.Println("==========================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Println("==========================")
		fmt.Print("Choose an option: ")

		// Membaca input dari user
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Menangani pilihan user
		switch choice {
		case "1":
			// Proses login
			user := login(reader, db)
			if user != nil {
				fmt.Printf("User '%s' (ID: %d) logged in!\n", user.Email, user.ID)
				// Menampilkan menu sesuai role
				if user.Role != "admin" {
					// Menu untuk customer
					showMenuCustomer(reader, db, user)
				} else {
					// Menu untuk admin
					showMenuAdmin(reader, db, user)
				}
				// Tambahkan fitur utama setelah login di sini
			} else {
				fmt.Println("❌ Login failed.")
			}
		case "2":
			// Proses registrasi
			user := register(reader, db)
			if user != nil {
				fmt.Printf("User '%s' (ID: %d) registered!\n", user.Email, user.ID)
			} else {
				fmt.Println("❌ Register failed.")
			}
		case "0":
			// Keluar dari program
			fmt.Println("Exiting...")
			return
		default:
			// Pilihan tidak valid
			fmt.Println("Invalid option.")
		}
	}
}

// Login menangani proses login user berdasarkan email dan password
func login(reader *bufio.Reader, db *sql.DB) *entity.User {
	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')

	// Menghilangkan spasi/tanda newline
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	// Validasi input
	if email == "" || password == "" {
		fmt.Println("Email and password cannot be empty.")
		return nil
	}

	// Memanggil fungsi login dari handler
	user, err := handler.LoginUser(db, email, password)
	if err != nil {
		fmt.Println("Login error:", err)
		return nil
	}
	fmt.Println("✅ Login successful.")
	return user
}

// register menangani proses registrasi user baru
func register(reader *bufio.Reader, db *sql.DB) *entity.User {
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	fmt.Print("Enter role (admin / customer): ")
	role, _ := reader.ReadString('\n')

	// Menghilangkan spasi/tanda newline
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	role = strings.TrimSpace(role)

	// Validasi input awal
	if email == "" || password == "" || role == "" {
		fmt.Println("All fields are required.")
		return nil
	}

	// Input tambahan
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter phone number: ")
	phone, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)

	// Registrasi akun user
	user, err := handler.RegisterUser(db, email, password, role)
	if err != nil {
		fmt.Println("Register error: ", err)
		return nil
	}

	// Menyimpan detail user tambahan
	err = handler.RegisterUserDetails(db, user.ID, name, phone)
	if err != nil {
		fmt.Println("❌ Failed to register user details:", err)
		return nil
	}

	fmt.Println("✅ Register successful.")
	return user
}
