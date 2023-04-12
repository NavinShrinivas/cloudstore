package database

import (
	"gorm.io/gorm"
)

// -----Databse Models-----
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey" `
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	UserType string `json:"usertype" gorm:"type:ENUM('buyer','seller','admin')"`
	Address  string `json:"address"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType string `json:"usertype"`
	Address  string `json:"address"`
}
