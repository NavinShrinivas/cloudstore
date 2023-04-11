package main

import (
	"net/http"
	"time"

	authentication "userhandle/authentication"
	database "userhandle/database"
	routes "userhandle/routes"

	"github.com/gin-gonic/gin"
	log "github.com/urishabh12/colored_log"
)

func main() {
	log.Println("Starting user handle services...")

	database.InitDatabaseVaraiables()
	authentication.InitAuthVariables()

	r := gin.Default()
	routes.RouteHandler(r)

	s := &http.Server{
		Addr:         ":5001",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("User handle services started.")
	log.Println("Listening on port 5001.")
	s.ListenAndServe()
}
