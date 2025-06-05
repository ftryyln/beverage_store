package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	// Try to load .env explicitly relative to project root
	projectRoot := filepath.Join("..", ".env") // adjust this if test runs from /handler folder

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// .env does not exist in current folder, try loading from project root
		err := godotenv.Load(projectRoot)
		if err != nil {
			log.Println("Warning: Could not load .env file from project root:", err)
		}
	} else {
		// load .env from current folder (normal case)
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Could not load .env file from current folder:", err)
		}
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed connection to the database!", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database cannot be accessed!", err)
	}

	return db
}
