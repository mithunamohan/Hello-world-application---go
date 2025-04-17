package models

// Greetings represents a document in the "greetings" collection
type Greetings struct {
	Greetings string `bson:"greetings"` // Maps the field to MongoDB
}
