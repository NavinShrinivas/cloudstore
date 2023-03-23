package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

func JWTAuthCheck(c *gin.Context) {
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
		return []byte(""), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "We are facing issues on our side, please try again later.",
		})
		log.Println("[WARN] Error on parsing token from header!!!")
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
func main() {
	log.Println("Starting user handle services...")
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user handle services are live.",
		})
	})

	r.Use(JWTAuthCheck) //[NOTE] All endpoints below this need auth (should have `Token` header in the request)

	s := &http.Server{
		Addr:           ":5001",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on port 5001.")
	s.ListenAndServe()
}
