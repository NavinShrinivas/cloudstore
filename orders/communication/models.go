package communication

import (
	"github.com/golang-jwt/jwt"
)

// -----Communication models-----

type LoginClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserType string `json:"usertype"`
	// type : buyer|seller|admin
	jwt.StandardClaims
}
type AuthResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"claims"`
}
