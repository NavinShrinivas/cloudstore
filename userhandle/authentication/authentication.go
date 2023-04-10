package authentication

import (
	// "fmt"
	"net/http"
	"os"
	"strconv"
	"userhandle/communication"
	"userhandle/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
)

func InitAuthVariables() {
	secret, err := godotenv.Read()
	if err != nil {
		log.Panic("Error reading .env file")
	}
	if secret["JWT_SECRET"] == "" {
		log.Panic("JWT_SECRET not set in .env file")
	}

	if secret["JWT_EXPIRE_TIME"] == "" {
		log.Panic("JWT_EXPIRE_TIME not set in .env file")
	}
	_, err = strconv.Atoi(secret["JWT_EXPIRE_TIME"])
	if err != nil {
		log.Panic("JWT_EXPIRE_TIME is not a number")
	}
	if secret["JWT_EXPIRE_TIME"] == "0" {
		log.Panic("JWT_EXPIRE_TIME cannot be 0")
	}

	os.Setenv("JWT_SECRET", secret["JWT_SECRET"])
	os.Setenv("JWT_EXPIRE_TIME", secret["JWT_EXPIRE_TIME"])
}

func JWTAuthCheck(c *gin.Context) {

	//[DONE] We can improve security by doing a database check after getting the claims
	parser_struct := jwt.Parser{
		UseJSONNumber:        true,  //Force number to be raw numbers and not strings
		SkipClaimsValidation: false, //Forces password validation
	}

	// get the token from cookie
	current_token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}

	log.Println("Token: ", current_token)

	//The second parameter is a callback function that Parse function executes
	claims := jwt.MapClaims{}
	token, err := parser_struct.ParseWithClaims(string(current_token[0]), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		log.Println("[WARN] Something going on with authentication!!!", err)
		c.Abort()
		return
	}
	if token.Valid {
		//Auth is succesfull
		//Do a quick DB lookup to actually see users existence
		user_query := database.User{
			Username: claims["username"].(string),
		}
		status, _ := database.GetUserRecord(user_query)
		if status {
			c.Next()
		} else {
			c.JSON(http.StatusProxyAuthRequired, gin.H{
				"status":  false,
				"message": "Auth failed, cannot find user records!",
			})
			log.Panic("[PANIC] JWT SECRET LEAKED!!!!!!!")
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

func GenerateAuthToken(user_record database.User) (string, int, bool, string, *communication.LoginClaims, int) {

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	expireTime, err := strconv.ParseInt(os.Getenv("JWT_EXPIRE_TIME"), 10, 64)
	if err != nil {
		log.Println("[WARN] Error parsing JWT_EXPIRE_TIME, using default value of 3600")
		expireTime = 3600
	}
	if expireTime <= 0 {
		log.Println("[WARN] Invalid JWT_EXPIRE_TIME, using default value of 3600")
		expireTime = 3600
	}

	// Create the Claims
	claims := communication.LoginClaims{
		Username: user_record.Username,
		// DOB:      user_record.DOB,
		UserType: user_record.UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "userservice",
			ExpiresAt: expireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "something went wrong on our side, please try again later.", http.StatusInternalServerError, false, "", nil, 0
	}
	return "Logged in succesfully!", http.StatusOK, true, ss, &claims, int(expireTime)
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
