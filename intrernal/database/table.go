package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"size:255"` // set field size to 255
}
