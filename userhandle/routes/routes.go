package routes

import (
	"net/http"
	"userhandle/authentication"
	"userhandle/communication"
	database "userhandle/database"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func AccountRouter(accountRoutes *gin.RouterGroup, r *gin.Engine) bool {

	accountRoutes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User handle API is running",
		})
	})

	accountRoutes.POST("/register", func(c *gin.Context) {
		var new_user database.User
		err := c.BindJSON(&new_user)
		if err != nil {
			log.Println("Error in binding json", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status := database.InsertUserRecord(new_user)
		c.JSON(httpstatus, gin.H{
			"status":  status,
			"message": message,
		})
	})

	accountRoutes.POST("/login", func(c *gin.Context) {
		var LoginUser communication.LoginRequest
		err := c.BindJSON(&LoginUser)
		if err != nil {
			log.Println("Error in binding json", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}

		message, httpstatus, status, user_record := database.CheckUserRecord(LoginUser)
		if status {
			message, httpstatus, status, token, claims, lifeTime := authentication.GenerateAuthToken(*user_record)
			if status {
				c.SetCookie("token", token, lifeTime, "/", "localhost", false, true)
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

	accountRoutes.POST("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Logged out successfully",
		})
	})

	accountRoutes.Use(authentication.JWTAuthCheck)

	accountRoutes.GET("/authcheck", func(c *gin.Context) {
		claims := authentication.GetClaimsInfo(c)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Valid user and token!",
			"claims":  claims,
		})
	})

	accountRoutes.POST("/authcheck", func(c *gin.Context) {
		claims := authentication.GetClaimsInfo(c)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Valid user and token!",
			"claims":  claims,
		})
	})

	accountRoutes.GET("/info", func(c *gin.Context) {
		claims := authentication.GetClaimsInfo(c)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		user_query := database.User{
			Username: claims["username"].(string),
		}

		if user_query.Username == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}

		status, user_record := database.GetUserRecord(user_query)
		if !status {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}

		c.JSON(http.StatusFound, gin.H{
			"status":  true,
			"message": "User record found",
			"record":  user_record,
		})
	})

	accountRoutes.PUT("/update", func(c *gin.Context) {
		var UpdateUser communication.EditRequest
		err := c.BindJSON(&UpdateUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		claims := authentication.GetClaimsInfo(c)
		message, httpstatus, status := database.UpdateUserRecord(UpdateUser, claims)
		c.JSON(httpstatus, gin.H{
			"status":  status,
			"message": message,
		})
	})

	accountRoutes.DELETE("/delete", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		var DeleteUser communication.DeleteRequest
		err := c.BindJSON(&DeleteUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status := database.DeleteUserRecord(DeleteUser)
		c.JSON(httpstatus, gin.H{
			"status":  status,
			"message": message,
		})
	})
	return false
}

func RouteHandler(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User handle services are live.",
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
}
