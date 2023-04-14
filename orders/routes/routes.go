package routes

import (
	"log"
	"net/http"
	"orders/authentication"
	"orders/communication"
	"orders/database"

	"github.com/gin-gonic/gin"
)

func OrderRouter(orderRoutes *gin.RouterGroup, r *gin.Engine) bool {

	orderRoutes.GET("/status", func(c *gin.Context) {
		database.GetDatabaseConnection()
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Order services API is running.",
		})
	})

	orderRoutes.Use(authentication.CheckUserAuthMiddleware)

	orderRoutes.GET("/fetch", func(c *gin.Context) {

		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		var b database.GetOrderRequest
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. Please check the request body.",
			})
			log.Println("[WARN] Invalid request body for order fetch")
			return
		}

		message, httpstatus, status, order := database.GetOrder(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"username": claims.Username,
				"order":    order,
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

	orderRoutes.GET("/all", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		message, httpstatus, status, orders := database.GetAllOrders(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"username": claims.Username,
				"orders":   orders,
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

	orderRoutes.POST("/create", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		if claims.UserType != "admin" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer.",
			})
			log.Println("[WARN] Normal user trying to gain access to orders, leak of orders endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.CreateOrderRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, order := database.InsertOrder(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"order":   order,
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

	orderRoutes.PUT("/update", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		var b communication.UpdateOrderRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, order := database.UpdateOrder(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"order":   order,
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

	orderRoutes.DELETE("/delete", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer.",
			})
			log.Println("[WARN] Normal user trying to gain access to orders, leak of orders endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.DeleteOrderRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, order := database.DeleteOrder(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"order":   order,
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
			"message": "Product services are live.",
		})
	})

	apiRoutes := r.Group("/api")
	{
		orderRoutes := apiRoutes.Group("/orders")
		{
			OrderRouter(orderRoutes, r)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "404 page not found",
		})
	})
}
