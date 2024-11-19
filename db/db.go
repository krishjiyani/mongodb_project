package db

import (
	"context"
	"crypto/tls"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore is the implementation of the Store interface for MongoDB.
type MongoStore struct {
	Client *mongo.Client
}

// NewMongoStore creates a new MongoStore and connects to the MongoDB server.
func NewMongoStore(uri, dbName string) (*MongoStore, error) {
	tlsConfig := &tls.Config{}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).
		SetServerAPIOptions(serverAPI).
		SetTLSConfig(tlsConfig)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// Ping MongoDB to confirm connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return &MongoStore{
		Client: client,
	}, nil
}