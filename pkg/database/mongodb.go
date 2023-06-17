package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(connectionString, databaseName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(databaseName)

	return &MongoDB{
		client:   client,
		database: database,
	}, nil
}

func (db *MongoDB) Close() {
	err := db.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}


func (db *MongoDB) InsertDocument(collectionName string, document interface{}) error {
	collection := db.database.Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}
