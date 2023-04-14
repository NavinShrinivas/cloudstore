package main

import (
	"net/http"
	"os"
	"products/routes"
	"strings"
	"time"

<<<<<<< HEAD
=======
	envLoader "products/envLoader"

>>>>>>> d36864c9dcef24a31fa4c751e153cd1e7690570b
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {

	log.Println("Starting product services...")

	envLoader.CheckAndSetVariables()

	r := gin.Default()
	config := cors.DefaultConfig()
<<<<<<< HEAD
	config.AllowOrigins = []string{"http://localhost:3000"}
	authentication.InitAuthVariables()
	database.InitDatabaseVaraiables()
	config.AllowCredentials = true
	r.Use(cors.New(config))

=======
	config.AllowOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	config.AllowCredentials = true
	r.Use(cors.New(config))
>>>>>>> d36864c9dcef24a31fa4c751e153cd1e7690570b
	routes.RouteHandler(r)

	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("Product services started and running at " + s.Addr)
	s.ListenAndServe()
}
