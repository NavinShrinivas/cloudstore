package database

import (
	"gorm.io/gorm"
)

// -----Databse Models-----
type User struct {
	gorm.Model
	Username string `gorm:"primaryKey" json:"username"`
	// DOB      datatypes.Date `json:"dob"`
	Password string `json:"password"` //[TODO]Let's skip out on hashing for inital release
	UserType string `json:"usertype"`
	//type : buyer|seller|admin
}
