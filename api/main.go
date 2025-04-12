package main

import (
	"fmt"
	"go-project-deb/apiroutes"
	// "go-project-deb/api_routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiroutes.RegisterRoutes(router)

	fmt.Println("Example app listening on port 8080")
	router.Run() // default: ":8080"
}


// func RegisterRoutes(router *gin.Engine) {
// 	panic("unimplemented")
// }
