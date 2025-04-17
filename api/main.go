package main

import (
	"go-project-deb/config"
	"go-project-deb/apiroutes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize DB connection
	config.ConnectDB()

	// Setup router
	router := gin.Default()
	apiroutes.RegisterRoutes(router)

	log.Println("Example app listening on port 8080")
	router.Run(":8081")
}
