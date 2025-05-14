package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Load environment variables from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
	if client == nil {
		uri := os.Getenv("MONGO_URI")
		var err error
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
	}
	return client
}

// GetDatabase returns the MongoDB database
func GetDatabase(dbName string) *mongo.Database {
	return GetClient().Database(dbName)
}

// GetCollection returns a specific collection from the MongoDB database
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return GetDatabase(dbName).Collection(collectionName)
}
