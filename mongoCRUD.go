package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var err error

func connectToMongo(uri string, cred *options.Credential, MongoClient *mongo.Client) (*mongo.Client, context.Context, error) {
	MongoClient, err = mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(*cred))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal()
	}

	return MongoClient, ctx, err
}

func returnCollection(connection *mongo.Client, tableName string, collectionName string) *mongo.Collection {
	return connection.Database(tableName).Collection(collectionName)
}

func dropCollection(collection *mongo.Collection, ctx context.Context) {
	if err := collection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}

func insertOne(collection *mongo.Collection, ctx context.Context, document interface{}) *mongo.InsertOneResult {
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func findOne(collection *mongo.Collection, ctx context.Context, id interface{}) *mongo.SingleResult {
	result := collection.FindOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CloseClientDB(MongoClient *mongo.Client, ctx context.Context) {
	if MongoClient == nil {
		return
	}

	err := MongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}
