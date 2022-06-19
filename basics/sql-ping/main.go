package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbPass = ""
	dbHost = "localhost"
	dbPort = 3306
	dbName = "edspert"
)

func main() {
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s", dbUser, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("berhasil terkoneksi")
}
