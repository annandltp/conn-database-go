package main

import (
	"fmt"

	"edspert.id/golang/database/database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id          int64
	name, email string
}

// Exec for INSERT, CREATE, UPDATE, DROP, DELETE
func main() {
	db := database.ConnectDB()
	defer db.Close()

	// QUERY MULTIPLE ROW

	// query := "SELECT * FROM users"
	// rows, err := db.Query(query)
	// if err != nil {
	// 	panic(err)
	// }

	// var users []User

	// for rows.Next() {
	// 	var user User

	// 	err := rows.Scan(&user.id, &user.name, &user.email)
	// 	if err != nil {
	// 		log.Fatal("terjadi error", err)
	// 	}

	// 	users = append(users, user)
	// }

	// fmt.Println(users)

	// QUERY SINGLE ROW

	query := "SELECT * FROM users WHERE email = ?"
	email := "nizomsidiq@gmail.com"
	row := db.QueryRow(query, email)

	var user User

	err := row.Scan(&user.id, &user.name, &user.email)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
