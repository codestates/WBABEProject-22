package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/config"
)

var DB *mongo.Client

var ProductCollection *mongo.Collection
var OrderCollection *mongo.Collection
var ReviewCollection *mongo.Collection

func ConnectDB(cfg *config.Config) {
	cf := cfg.DB
	uri := cf["host"]
	databaseName := cf["name"]
	
	DB = getDatabase(uri)
	
	ProductCollection = GetCollection(DB, databaseName, "products")
	OrderCollection = GetCollection(DB, databaseName, "orders")
	ReviewCollection = GetCollection(DB, databaseName, "reviews")
	
	// Product codes should be unique.
	_, err := ProductCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic(err)
	}
}

func getDatabase(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}

func GetCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
