package communication

import (
	"github.com/golang-jwt/jwt"
)

// -----Communication models-----
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserType string `json:"usertype"`
	// type : buyer|seller|admin
	jwt.StandardClaims
}

type LoginResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"Claims"`
	User    UserInfo    `json:"user"`
}

type EditRequest struct {
	Username string `json:"username"` // compulsory
	Password string `json:"password"` // compulsory
	// Optional fields
	NewUsername string `json:"newusername"`
	NewPassword string `json:"newpassword"`
	NewName     string `json:"newname"`
	NewEmail    string `json:"newemail"`
	NewPhone    string `json:"newphone"`
	NewAddress  string `json:"newaddress"`
	NewUserType string `json:"newusertype"`
	// type : buyer|seller|admin
}

type DeleteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType string `json:"usertype"`
	Address  string `json:"address"`
}
