package apiroutes

import (
    "math/rand"
	"net/http"
	"time"
    "github.com/gin-gonic/gin"
)
var greetingsList = []string{
	"Hello World!",
	"Bonjour monde!",
	"Hallo Welt!",
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine) {
	// GET /
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// GET /greetings
	router.GET("/greetings", func(c *gin.Context) {
		c.JSON(http.StatusOK, greetingsList)
	})

	// GET /visitors
	router.GET("/visitors", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		visitorCount := rand.Intn(1000)
		c.JSON(http.StatusOK, gin.H{"visitorCount": visitorCount})
	})
}
