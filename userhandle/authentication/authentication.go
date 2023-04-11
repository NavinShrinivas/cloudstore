package authentication

import (
	// "fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"userhandle/communication"
	"userhandle/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
)

func InitAuthVariables() {

	envRequired := []string{"JWT_SECRET", "JWT_LIFETIME"}

	_, err := os.Stat(".env")
	if err == nil {
		secret, err := godotenv.Read()
		if err != nil {
			log.Panic("Error reading .env file")
		}

		for _, key := range envRequired {
			if secret[key] != "" {
				os.Setenv(key, secret[key])
			}
		}
	}

	for _, key := range envRequired {
		if os.Getenv(key) == "" {
			log.Panic("Environment variable " + key + " not set")
		}
	}

	lifeTime, err := strconv.ParseInt(os.Getenv("JWT_LIFETIME"), 10, 64)
	if err != nil {
		log.Panic("Error parsing JWT_LIFETIME")
	}
	if lifeTime <= 0 {
		log.Panic("Invalid JWT_LIFETIME")
	}
}

func GenerateAuthToken(user_record database.User) (string, int, bool, string, *communication.LoginClaims, int) {

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	lifeTime, _ := strconv.ParseInt(os.Getenv("JWT_LIFETIME"), 10, 64)

	// Create the Claims
	claims := communication.LoginClaims{
		Name:     user_record.Name,
		Username: user_record.Username,
		UserType: user_record.UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "userservice",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: lifeTime + time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "Something went wrong on our side, please try again later.", http.StatusInternalServerError, false, "", nil, 0
	}
	return "Logged in succesfully!", http.StatusOK, true, ss, &claims, int(lifeTime)
}

func JWTAuthCheck(c *gin.Context) {

	parser_struct := jwt.Parser{
		UseJSONNumber:        true,  //Force number to be raw numbers and not strings
		SkipClaimsValidation: false, //Forces password validation
	}

	current_token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}

	claims := jwt.MapClaims{}
	token, err := parser_struct.ParseWithClaims(string(current_token), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "Auth failed! Invalid JWT token.",
		})
		log.Println("[WARN] Something going on with authentication!!!", err)
		c.Abort()
		return
	}
	if token.Valid {
		user_query := database.User{
			Username: claims["username"].(string),
		}
		status, _ := database.GetUserRecord(user_query)
		if status {
			c.Next()
		} else {
			c.JSON(http.StatusProxyAuthRequired, gin.H{
				"status":  false,
				"message": "Auth failed! Invalid JWT token.",
			})
			log.Panic("[PANIC] JWT SECRET BREACHED!!! Username:"+claims["username"].(string), err)
		}
	} else {
		c.JSON(http.StatusProxyAuthRequired, gin.H{
			"status":  false,
			"message": "Auth failed! Invalid JWT token.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}
}

func GetClaimsInfo(c *gin.Context) map[string]interface{} {
	//[TODO]Maybe possible to merge with middleware auth check
	//[TODO]Maybe can inject to request headers in middleware only
	parser_struct := jwt.Parser{
		UseJSONNumber:        true,  //Force number to be raw numbers and not strings
		SkipClaimsValidation: false, //Forces password validation
	}

	current_token, err := c.Cookie("token")
	if err != nil {
		log.Panic("Auth token not found in cookie, will report to admins.")
	}

	claims := jwt.MapClaims{}
	token, _ := parser_struct.ParseWithClaims(string(current_token[0]), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if token.Valid {
		//Above if condition is redundant
		return claims //Is a hashmap k-v pair
	}
	log.Panic("Someone is messing with memory stuff!!!")
	return nil
}
