package database

import (
	"context"
	"fmt"
	"short-url/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(short string) (models.Data, error) {

	client, ctx, collection := Connect()
	defer client.Disconnect(ctx)

	filter := bson.M{"short": short}

	var shortURL models.Data
	err := collection.FindOne(context.Background(), filter).Decode(&shortURL)
	if err != nil {
		fmt.Println(err)
	}
	return shortURL, err
}
