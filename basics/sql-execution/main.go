package main

import (
	"context"
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

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s", dbUser, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

// Exec for INSERT, CREATE, UPDATE, DROP, DELETE
func main() {
	db := ConnectDB()
	defer db.Close()

	ctx := context.Background()

	// sql := `CREATE TABLE users (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	name VARCHAR(255) NOT NULL,
	// 	email VARCHAR(255) NOT NULL UNIQUE
	// )`
	// _, err := db.ExecContext(ctx, sql)
	// if err != nil {
	// 	panic(err)
	// }

	sql := "INSERT INTO users(name, email) VALUES (?, ?)"

	var (
		name  = "nizom"
		email = "nizomsidiq@gmail.com"
	)
	result, err := db.ExecContext(ctx, sql, name, email)
	if err != nil {
		panic(err)
	}

	var user User
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.id = id
	user.name = name
	user.email = email

	fmt.Println(user)
}

type User struct {
	id          int64
	name, email string
}
