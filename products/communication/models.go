package communication

import (
	"github.com/golang-jwt/jwt"
)

type LoginClaims struct {
	Username string `json:"username"`
	UserType string `json:"usertype"`
	//type : buyer|seller|admin
	jwt.StandardClaims
}

type AuthResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"claims"`
}

type CreateProductRequest struct {
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Limit        int     `json:"limit"`
	Manufacturer string  `json:"manufacturer"`
}

type EditProductRequest struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Limit        int     `json:"limit"`
	Manufacturer string  `json:"manufacturer"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}
