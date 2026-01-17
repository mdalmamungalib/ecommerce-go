package database

import (
	"context"
	"log"
	"time"

	"ecommerce/constant" // Imports your constant package for URI and DB name

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientInstance is a global variable to hold the MongoDB connection
var ClientInstance *mongo.Client

// ConnectDb initializes the connection to MongoDB
func ConnectDb() {
	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options using the URI from your constant package
	clientOptions := options.Client().ApplyURI(constant.MDBUri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB: ", err)
	}

	// Check the connection (Ping)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Database is not responding: ", err)
	}

	log.Println("✅ Successfully connected to MongoDB!")
	ClientInstance = client
}

// OpenCollection is a helper function to access a specific collection
func OpenCollection(collectionName string) *mongo.Collection {
	if ClientInstance == nil {
		log.Fatal("❌ Database client is not initialized. Call ConnectDb first.")
	}
	return ClientInstance.Database(constant.Database).Collection(collectionName)
}