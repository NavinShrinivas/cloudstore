package main

import (
	// "fmt"
	log "github.com/urishabh12/colored_log"
	"github.com/gin-gonic/gin"
   "net/http"
	"products/authentication"
   "time"
)

func main() {
	log.Println("Starting product services...")
	r := gin.Default()
	r.Use(authentication.CheckUserAuthMiddleware) //[NOTE] All endpoints below this need auth (should have `Token` header in the request)
	//If a person has another persons valid JWT token they can wrek havoc
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "product services are live.",
		})
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
