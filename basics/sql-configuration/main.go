package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// jumlah koneksi yang bisa nganggur (idle)
	db.SetMaxIdleConns(2)

	// idle (nganggur) selama 1 jam, setelah itu koneksi akan diputus
	db.SetConnMaxIdleTime(1 * time.Hour)

	// jumlah koneksi yang bisa dibuka, FYI: database punya jumlah maksimal koneksi
	db.SetMaxOpenConns(5)

	// jumlah maksimal lamanya koneksi
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}
