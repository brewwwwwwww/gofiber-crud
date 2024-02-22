package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:not null`
	Price       float64 `gorm:not null`
	Description string
}

type User struct {
	gorm.Model
	FirstName string `gorm:not null`
	LastName  string `gorm:not null`
	Email     string `gorm:unique; not null`
	Password  string `gorm:not null`
}
