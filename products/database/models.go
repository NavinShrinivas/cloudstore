package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string  `json:"name"`
	Limit          int     `json:"limit"`
	Username       string  `json:"username"`
	Price          float32 `json:"price"`
	Avgrating      float32 `json:"avgrating"`
	Numberofrating int     `json:"numberofratings"`
	Manufacturer   string  `json:"manufacturer"`
}
