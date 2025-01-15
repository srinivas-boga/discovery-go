package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB connects to MongoDB and returns a client
func connectToMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetMongoDBCollection returns a collection from a MongoDB database
func GetMongoDBCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

// GetMongoDBClient returns a MongoDB client
func GetMongoDBClient() *mongo.Client {
	return connectToMongoDB()
}
