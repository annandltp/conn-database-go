package main

import (
	"fmt"

	"edspert.id/golang/database/gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser = "root"
	dbPass = ""
	dbHost = "localhost"
	dbPort = 3306
	dbName = "edspert"
)

func main() {
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Profile{})

	// user := &models.User{
	// 	FirstName: "Nizom",
	// 	LastName:  "Sidiq",
	// 	Email:     "nizomsidiq@gmail.com",
	// }
	// db.Create(user)

	var user models.User
	findId := 1

	db.First(&user, "id = ?", findId)

	fmt.Println("Get single User")
	fmt.Println("====")
	fmt.Println("id", user.ID)
	fmt.Println("firstname", user.FirstName)
	fmt.Println("lastname", user.LastName)

	fmt.Println("Get multiple User")
	fmt.Println("====")

	var users []models.User
	db.Find(&users)

	for _, user := range users {
		fmt.Println("id", user.ID)
		fmt.Println("firstname", user.FirstName)
		fmt.Println("lastname", user.LastName)
	}

	// db.Model(users[1]).Update("LastName", "Doe")

	// profile := &models.Profile{
	// 	Address:     "Yogyakarta",
	// 	PhoneNumber: "08123456789",
	// 	UserID:      users[0].ID,
	// }

	// db.Create(profile)

	var updatedUser models.User
	db.Model(&updatedUser).Association("Profile")
	db.First(&updatedUser, users[0].ID)

	fmt.Println(updatedUser.Profile)
}
