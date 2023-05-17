package database

import (
	"context"
	"fmt"
	"short-url/backend/models"
)

func Register(short string, url string) error {

	client, ctx, collection := Connect()
	defer client.Disconnect(ctx)

	filter := models.Data{Short: short, Url: url}

	_, err := collection.InsertOne(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
