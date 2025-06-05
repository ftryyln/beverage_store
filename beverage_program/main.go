package main

import (
	"beverage_program/cli"
	"beverage_program/config"
	"bufio"
	"os"
)

func main() {
	// Inisialisasi koneksi database
	db := config.InitDB()
	// Pastikan koneksi database ditutup saat program selesai
	defer db.Close()

	// Membuat buffered reader untuk input dari user di command line
	reader := bufio.NewReader(os.Stdin)

	// Menampilkan menu utama CLI dan mulai interaksi dengan user
	cli.ShowMenu(reader, db)
}
