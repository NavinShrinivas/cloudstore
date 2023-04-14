package routes

import (
	"log"
	"net/http"
	"reviews/authentication"
	"reviews/communication"
	"reviews/database"

	"github.com/gin-gonic/gin"
)

func ReviewRouter(reviewRoutes *gin.RouterGroup, r *gin.Engine) bool {

	reviewRoutes.GET("/status", func(c *gin.Context) {
		database.GetDatabaseConnection()
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Review services API is running.",
		})
	})

	reviewRoutes.Use(authentication.CheckUserAuthMiddleware)

	reviewRoutes.GET("/fetch", func(c *gin.Context) {

		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		var b database.GetReviewRequest
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. Please check the request body.",
			})
			log.Println("[WARN] Invalid request body for review fetch")
			return
		}

		message, httpstatus, status, review := database.GetReview(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"username": claims.Username,
				"review":   review,
			})
			return
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
			return
		}
	})

	reviewRoutes.GET("/all", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		message, httpstatus, status, reviews := database.GetAllReviews(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"username": claims.Username,
				"reviews":  reviews,
			})
			return
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
			return
		}
	})

	reviewRoutes.POST("/create", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		if claims.UserType != "admin" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer.",
			})
			log.Println("[WARN] Normal user trying to gain access to reviews, leak of reviews endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.CreateReviewRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, review := database.InsertReview(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"review":  review,
			})
			return
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
			return
		}
	})

	reviewRoutes.PUT("/update", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		if claims.UserType != "admin" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer.",
			})
			log.Println("[WARN] Normal user trying to gain access to reviews, leak of reviews endpoints. Possible DDOS attempt.")
			return
		}

		var b communication.UpdateReviewRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, review := database.UpdateReview(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"review":  review,
			})
			return
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
			return
		}
	})

	reviewRoutes.DELETE("/delete", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer.",
			})
			log.Println("[WARN] Normal user trying to gain access to reviews, leak of reviews endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.DeleteReviewRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, review := database.DeleteReview(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"review":  review,
			})
			return
		} else {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
			})
			return
		}
	})

	return false
}

func RouteHandler(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Review services are live.",
		})
	})

	apiRoutes := r.Group("/api")
	{
		reviewRoutes := apiRoutes.Group("/reviews")
		{
			ReviewRouter(reviewRoutes, r)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "404 page not found",
		})
	})
}
