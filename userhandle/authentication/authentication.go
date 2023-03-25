package authentication

import (
	// "fmt"
	"net/http"
	"userhandle/communication"
	"userhandle/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/urishabh12/colored_log"
)

func JWTAuthCheck(c *gin.Context) {
   //[TODO] We can improve security by doing a database check after getting the claims
	parser_struct := jwt.Parser{
		UseJSONNumber:        true,  //Force number to be raw numbers and not strings
		SkipClaimsValidation: false, //Forces password validation
	}
	current_token_header := c.Request.Header["Token"]
	if len(current_token_header) == 0 {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in header, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}

	//The second parameter is a callback function that Parse function executes
	token, err := parser_struct.Parse(current_token_header[0], func(token *jwt.Token) (interface{}, error) {
		return []byte("shoouldbekeptsecret"), nil
		//[TODO]Should move the password to a global hidden config file
	})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
         "status" : false,
			"message": err.Error(),
		})
		log.Println("[WARN] Something going on with authentication!!!",err)
		c.Abort()
		return
	}
	if token.Valid {
		//Auth is succesfull
		c.Next()
	} else {
		c.JSON(http.StatusProxyAuthRequired, gin.H{
			"message": "Auth failed! Invalid JWT token.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}
}

func GenerateAuthToken(user_record database.User) (string,int,bool,string,*communication.LoginClaims){
	mySigningKey := []byte("shoouldbekeptsecret")
	// Create the Claims
	claims := communication.LoginClaims{
		user_record.Username,
      user_record.DOB,
      user_record.UserType,
		jwt.StandardClaims{
			Issuer:    "userservice",
		},
      //[TODO] No expiry time, should implement it.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
   if err!=nil{
      return "something went wrong on our side, please try again later.",http.StatusInternalServerError,false,"",nil
   }
   return "Logged in succesfully!",http.StatusOK,true,ss,&claims
}


func GetClaimsInfo(c *gin.Context) map[string]interface{}{
   //[TODO]Maybe possible to merge with middleware auth check 
   //[TODO]Maybe can inject to request headers in middleware only
	parser_struct := jwt.Parser{
		UseJSONNumber:        true,  //Force number to be raw numbers and not strings
		SkipClaimsValidation: false, //Forces password validation
	}
	current_token_header := c.Request.Header["Token"]
   claims := jwt.MapClaims{}
	token, _ := parser_struct.ParseWithClaims(current_token_header[0], claims,func(token *jwt.Token) (interface{}, error) {
		return []byte("shoouldbekeptsecret"), nil
		//[TODO]Should move the password to a global hidden config file
	})
   if token.Valid{
      //Above if condition is redundant 
      return claims//Is a hashmap k-v pair
   }
   log.Panic("Someeone is messing with memory stuff!!")
   return nil
}
