package apiroutes

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go-project-deb/config" // Import the config package
	"go-project-deb/models" // Import the models package
)

// List of greetings
// var greetingsList = []string{
// 	"Hello World!",
// 	"Bonjour monde!",
// 	"Hallo Welt!",
// }

// RegisterRoutes sets up all the routes with CORS middleware and handlers
func RegisterRoutes(router *gin.Engine) {
	// Apply CORS middleware only to this router group
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with actual frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Route handler for "/"
	router.GET("/", func(c *gin.Context) {
		// Handler for the home route
		c.String(http.StatusOK, "Hello World! Welcome to the Go API!")
	})

	// Route handler for "/greetings"
	router.GET("/greetings", func(c *gin.Context) {
		collection := config.GetCollection("greetings")

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch greetings"})
			return
		}
		defer cursor.Close(context.TODO())

		var greetings []models.Greetings
		if err := cursor.All(context.TODO(), &greetings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode greetings"})
			return
		}

		var greetingsList []string
		for _, g := range greetings {
			greetingsList = append(greetingsList, g.Greetings)
		}

		c.JSON(http.StatusOK, greetingsList)
	})

	// Route handler for "/visitors"
	router.GET("/visitors", func(c *gin.Context) {
		collection := config.GetCollection("visitors")
		filter := bson.M{"name": "latest"}

		var visitor models.Visitors
		err := collection.FindOne(context.TODO(), filter).Decode(&visitor)
		if err == nil {
			// Document exists - increment count
			visitor.VisitorCount++
			update := bson.M{"$set": bson.M{"visitorcount": visitor.VisitorCount}}
			_, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update visitor count"})
				return
			}
			c.JSON(http.StatusOK, visitor.VisitorCount)
			return
		}

		// If document doesn't exist, create one
		_, err = collection.InsertOne(context.TODO(), bson.M{
			"name":         "latest",
			"visitorcount": 1,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create visitor document"})
			return
		}

		c.JSON(http.StatusOK, 1)
	})
}