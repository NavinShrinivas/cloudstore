package main

import (
	"net/http"
	"time"

	authentication "orders/authentication"
	database "userhandle/database"
	routes "userhandle/routes"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {
	log.Println("Starting order services...")

	database.InitDatabaseVaraiables()
	authentication.InitAuthVariables()

	r := gin.Default()
	routes.RouteHandler(r)

	s := &http.Server{
		Addr:         ":5003",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("Order services started.")
	log.Println("Listening on port 5003.")
	s.ListenAndServe()
}
