package main

import (
	"net/http"
	"os"
	"time"

	envLoader "userhandle/envLoader"
	routes "userhandle/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {

	log.Println("Starting user handle services...")

	envLoader.CheckAndSetVariables()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CORS_ORIGIN")}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	routes.RouteHandler(r)

	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("User handle services started and running at " + s.Addr)
	s.ListenAndServe()
}
