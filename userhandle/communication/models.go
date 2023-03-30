package communication

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/datatypes"
)

// -----Communication models-----
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginClaims struct {
	Username string         `json:"username"`
	DOB      datatypes.Date `json:"dob"`
	UserType string         `json:"usertype"`
	//type : buyer|seller|admin
	jwt.StandardClaims
}

type LoginResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"Claims"`
	Token   string      `json:"token"`
}

type EditRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

