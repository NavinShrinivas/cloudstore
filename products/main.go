package main

import (
	// "fmt"
	"net/http"
	"products/authentication"
	"products/communication"
	"products/database"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {
	log.Println("Starting product services...")
	r := gin.Default()
	r.Use(authentication.CheckUserAuthMiddleware) //[NOTE] All endpoints below this need auth (should have `Token` header in the request)
	//If a person has another persons valid JWT token they can wrek havoc
	r.GET("/productstatus", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "product services are live.",
		})
	})

   //Business logic : Will move to insert to database only if the user is of type seller
	r.POST("/product", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
      if claims.UserType != "seller"{
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
		message, httpstatus, status, productid := database.SetProduct(b, claims)
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

	r.PUT("/product", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
      if claims.UserType != "seller"{
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller.",
			})
         log.Println("[WARN] Normal user trying to gain access to products, leak of protected endpoints. Possible DDOS attempt.")
			return
      }
		var b communication.EditProductRequest
		err := c.BindJSON(&b)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request",
			})
			return
		}
		message, httpstatus, status:= database.EditProduct(b, claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":    status,
				"message":   message,
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

	r.GET("/product", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
      if claims.UserType != "buyer"{
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for buyer accounts.",
			})
			return
      }
		message, httpstatus, status, products:= database.GetAllProducts()
		if status {
			c.JSON(httpstatus, gin.H{
				"status":    status,
				"message":   message,
            "items": products,
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

	r.GET("/sellerproduct", func(c *gin.Context) {
		wrappedclaims := authentication.GetClaims(c)
		claims := wrappedclaims.Claims
      if claims.UserType != "seller"{
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "invalid request. This endpoint is valid only for seller accounts.",
			})
			return
      }
		message, httpstatus, status, products:= database.GetAllSellerProducts(claims)
		if status {
			c.JSON(httpstatus, gin.H{
				"status":    status,
				"message":   message,
            "username": claims.Username,
            "items": products,
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
	s := &http.Server{
		Addr:         ":5002",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on port 5002.")
	s.ListenAndServe()
}
