package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
    Client *mongo.Client
    DB     *mongo.Database
}

var Client *MongoDBClient

func ConnectMongoDBClient(connectionString, dbName string)  {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
    if err != nil {
		log.Fatal("Can't connect MongoDB file")
	}
    _, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    Client = &MongoDBClient{
        Client: client,
        DB:     client.Database(dbName),
    }
}


func (c *MongoDBClient) Close() error {
    return c.Client.Disconnect(context.Background())
}

func (c *MongoDBClient) GetCollection(collectionName string) *mongo.Collection {
    return c.DB.Collection(collectionName)
}
