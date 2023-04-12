package main

import (
	"net/http"
	"products/authentication"
	"products/database"
	"products/routes"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {
	log.Println("Starting product services...")
	r := gin.Default()

	authentication.InitAuthVariables()
	database.InitDatabaseVaraiables()

	routes.RouteHandler(r)

	s := &http.Server{
		Addr:         ":5002",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("Product handle services started.")
	log.Println("Listening on port 5002.")
	s.ListenAndServe()
}
