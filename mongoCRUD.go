package main

import (
	"Electric_Characteristics/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var mongoConnection, ctx, _ = connectToMongo(config.Uri, &config.Cred, config.MongoClient)
var fd = returnExcelData(config.FilePath, config.SheetIndex)

func connectToMongo(uri string, cred *options.Credential, MongoClient *mongo.Client) (*mongo.Client, context.Context, error) {

	var err error

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
		log.Fatal(err)
	}

	return MongoClient, ctx, err
}

func returnCollection(connection *mongo.Client, tableName string, collectionName string) *mongo.Collection {
	return connection.Database(tableName).Collection(collectionName)
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
	log.Fatal("Connection to MongoDB closed.")
}

func dropCollection(connection *mongo.Client, tableName string, collectionName string, ctx context.Context) {
	dataCollection := returnCollection(connection, tableName, collectionName)
	if err := dataCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}

func findDataById(connection *mongo.Client, tableName string, collectionName string, id interface{}) (*mongo.SingleResult, error) {
	dataCollection := returnCollection(connection, tableName, collectionName)
	result := dataCollection.FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		log.Fatal(result.Err())
		return nil, result.Err()
	}
	return result, nil
}

func insertData(connection *mongo.Client, tableName string, collectionName string, ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	dataCollection := returnCollection(connection, tableName, collectionName)
	res, err := dataCollection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}
