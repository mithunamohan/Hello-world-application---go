package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// replace URI if needed

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	db = client
	log.Println("MongoDB connected!")
}

func GetCollection(collectionName string) *mongo.Collection {
	if db == nil {
		log.Fatal("MongoDB not connected. Call ConnectDB() first.")
	}
	return db.Database("deb").Collection(collectionName)
}

