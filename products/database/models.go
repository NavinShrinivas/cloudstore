package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string  `json:"name"`
	Limit          int     `json:"limit"`
	SellerUsername string  `json:"seller_username"`
	Price          float32 `json:"price"`
	AvgRating      float32 `json:"avg_rating"`
	NumberOfRating int     `json:"number_of_ratings"`
	Manufacturer   string  `json:"manufacturer"`
}
