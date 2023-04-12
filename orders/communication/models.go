package communication

import (
	"github.com/golang-jwt/jwt"
)

<<<<<<< HEAD
// -----Communication models----
=======
// -----Communication models-----
>>>>>>> b1491b0 (order mservice starting)

type LoginClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserType string `json:"usertype"`
	// type : buyer|seller|admin
	jwt.StandardClaims
}
<<<<<<< HEAD

=======
>>>>>>> b1491b0 (order mservice starting)
type AuthResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"claims"`
}
