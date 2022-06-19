package main

import (
	"context"
	"fmt"
	"log"

	"edspert.id/golang/database/database"
)

func main() {
	db := database.ConnectDB()

	queries := database.New(db)

	ctx := context.Background()
	users, err := queries.GetUsers(ctx)
	if err != nil {
		log.Fatal("gagal", err)
	}

	fmt.Println(users)
}
