package database

import (
	"gorm.io/gorm"
)

// -----Databse Models-----
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey" `
	Name     string `json:"name"`
	Email    string `json:"email"    gorm:"unique"`
	Phone    string `json:"phone"    gorm:"unique"`
	Password string `json:"password"`
	UserType string `json:"usertype" gorm:"type:ENUM('buyer','seller','admin')"`
	Address  string `json:"address"`
}
