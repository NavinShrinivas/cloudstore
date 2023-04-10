package routes

import (
	"net/http"
	"time"
	"userhandle/authentication"
	"userhandle/communication"
	database "userhandle/database"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func AccountRouter(accountRoutes *gin.RouterGroup, r *gin.Engine) bool {
	accountRoutes.POST("/register", func(c *gin.Context) {
		var b database.User
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

	accountRoutes.POST("/login", func(c *gin.Context) {
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
			message, httpstatus, status, token, claims, expireTime := authentication.GenerateAuthToken(*user_record)
			if status {
				c.SetCookie("token", token, expireTime, "/", "localhost", false, true)
				c.JSON(httpstatus, communication.LoginResponse{
					Status:  status,
					Message: message,
					Claims:  *claims,
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

	accountRoutes.Use(authentication.JWTAuthCheck)

	accountRoutes.PUT("/user", func(c *gin.Context) {
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

	accountRoutes.GET("/authcheck", func(c *gin.Context) {
		claims := authentication.GetClaimsInfo(c)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Valid user and token!",
			"claims":  claims,
		})
	})

	accountRoutes.POST("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Logged out successfully",
		})
	})
	return false
}

func RouteHandler(r *gin.Engine) {

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "user handle services are live.",
		})
	})

	apiRoutes := r.Group("/api")
	{
		accountRoutes := apiRoutes.Group("/account")
		{
			AccountRouter(accountRoutes, r)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "404 page not found",
		})
	})

	s := &http.Server{
		Addr:         ":5001",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Println("User handle services are live.")
	log.Println("Listening on port 5001.")
	s.ListenAndServe()
}
