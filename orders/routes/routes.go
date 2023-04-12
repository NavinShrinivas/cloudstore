package routes

import (
	"net/http"
	"orders/authentication"

	"github.com/gin-gonic/gin"
)

func AccountRouter(accountRoutes *gin.RouterGroup, r *gin.Engine) bool {

	accountRoutes.Use(authentication.CheckUserAuthMiddleware)

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
		accountRoutes := apiRoutes.Group("/orders")
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
