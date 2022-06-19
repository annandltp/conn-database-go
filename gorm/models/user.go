package models

import "gorm.io/gorm"

type User struct { // users
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Profile   *Profile
}
