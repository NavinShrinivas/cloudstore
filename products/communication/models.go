package communication

import(
	"github.com/golang-jwt/jwt"
   "gorm.io/datatypes"

)

type LoginClaims struct {
	Username string         `json:"username"`
	DOB      datatypes.Date `json:"dob"`
	UserType string         `json:"usertype"`
	//type : buyer|seller|admin
	jwt.StandardClaims
}

type AuthResponse struct {
   Status bool `json:"status"`
   Message string `json:"message"`
   Claims LoginClaims `json:"claims"`
}
