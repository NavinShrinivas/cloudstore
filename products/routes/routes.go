package routes

import (
	"log"
	"net/http"
	"products/authentication"
	"products/communication"
	"products/database"

	"github.com/gin-gonic/gin"
)

func ProductRouter(productRoutes *gin.RouterGroup, r *gin.Engine) bool {

	productRoutes.GET("/status", func(c *gin.Context) {
		database.GetDatabaseConnection()
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Product services API is running.",
		})
	})

	productRoutes.GET("/all", func(c *gin.Context) {
		message, httpstatus, status, products := database.GetAllProducts()
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"products": products,
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

	productRoutes.POST("/fetch", func(c *gin.Context) {

		var b communication.FetchProductsRequest

		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}

		message, httpstatus, status, products := database.GetProducts(b.IDs)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"products": products,
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

	productRoutes.Use(authentication.CheckUserAuthMiddleware)

	productRoutes.POST("/create", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims

		if claims.UserType != "admin" && claims.UserType != "seller" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller.",
			})
			log.Println("[WARN] Normal user trying to gain access to products, leak of protected endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.CreateProductRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, productid := database.InsertProduct(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":    status,
				"message":   message,
				"productid": productid,
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

	productRoutes.PUT("/update", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "seller" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller.",
			})
			log.Println("[WARN] Normal user trying to gain access to products, leak of protected endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.UpdateProductRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, product := database.UpdateProduct(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"product": product,
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

	productRoutes.DELETE("/delete", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "seller" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller.",
			})
			log.Println("[WARN] Normal user trying to gain access to products, leak of protected endpoints. Possible DDOS attempt.")
			return
		}
		var b communication.DeleteProductRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status := database.DeleteProduct(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
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

	productRoutes.GET("/seller", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "seller" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller accounts.",
			})
			return
		}
		message, httpstatus, status, products := database.GetAllSellerProducts(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":   status,
				"message":  message,
				"username": claims.Username,
				"items":    products,
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

	productRoutes.POST("/rate", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
		if claims.UserType != "admin" && claims.UserType != "seller" && claims.UserType != "buyer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller and buyer accounts.",
			})
			return
		}
		var b communication.RateProductRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status, product := database.RateProduct(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":  status,
				"message": message,
				"product": product,
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
		productRoutes := apiRoutes.Group("/products")
		{
			ProductRouter(productRoutes, r)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "404 page not found",
		})
	})
}
