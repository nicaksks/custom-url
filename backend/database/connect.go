package database

import (
	"context"
	"fmt"
	"log"
	"short-url/backend/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseURL = utils.Database()
)

type mongoClient struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func Connect() (*mongo.Client, context.Context, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Database & Collection
	collection := client.Database("list").Collection("urls")

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database online!")
	return client, ctx, collection
}
