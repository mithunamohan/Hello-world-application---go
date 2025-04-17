package models

// Visitors represents a document in the "visitors" collection
type Visitors struct {
	Name         string `bson:"name"`          // maps to name field in MongoDB
	VisitorCount int    `bson:"visitorcount"`  // maps to visitorcount field
}
