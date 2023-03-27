package main

import (
	// "fmt"
	"net/http"
	"time"

	"userhandle/authentication"
	"userhandle/communication"
	database "userhandle/database"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {
	log.Println("Starting user handle services...")
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "user handle services are live.",
		})
	})
	r.POST("/register", func(c *gin.Context) {
		var b database.User //Feel free to create different communication structure for your API's
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status := database.PutUserRecords(b)
		c.JSON(httpstatus, gin.H{
			"status":  status,
			"message": message,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var b communication.LoginRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, user_record := database.CheckUserRecords(b)
		if status {
			message, httpstatus, status, token, claims := authentication.GenerateAuthToken(*user_record)
			if status {
				c.JSON(httpstatus, communication.LoginResponse{
					Status:  status,
					Message: message,
					Claims:  *claims,
					Token:   token,
				})
			} else {
				c.JSON(httpstatus, gin.H{
					"status":  status,
					"message": message,
				})
			}
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
		}
	})
	r.Use(authentication.JWTAuthCheck) //[NOTE] All endpoints below this need auth (should have `Token` header in the request) 
   //If a person has another persons valid JWT token they can wrek havoc
	r.POST("/edit", func(c *gin.Context) {
		var b communication.EditRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		claims := authentication.GetClaimsInfo(c)
		message, httpstatus, status := database.UpdateUserRecord(b, claims)
		c.JSON(httpstatus, gin.H{
			"status":  status,
			"message": message,
		})
	})
	r.GET("/auth", func(c *gin.Context) {
		claims := authentication.GetClaimsInfo(c)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
         "message": "Valid user and token!",
         "claims" : claims,
		})
	})
	s := &http.Server{
		Addr:         ":5001",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on port 5001.")
	s.ListenAndServe()
}
