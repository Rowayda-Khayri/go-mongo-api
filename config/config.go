package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection details
const (
	mongoURL      = "mongodb://localhost:27017"
	databaseName  = "organizations_management"
	collectionName = "users"
)

// GetMongoDBClient returns a MongoDB client
func GetMongoDBClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetUsersCollection returns the "users" collection
func GetUsersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
