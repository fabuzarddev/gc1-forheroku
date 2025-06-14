package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// LoadEnv: Baca file .env sekali di awal program
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, pakai environment bawaan OS")
	}
}

// DBInit: Return koneksi DB sudah siap pakai
func DBInit() *sql.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Build connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}
	return db
}
