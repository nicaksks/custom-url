package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Delete(short string) error {

	client, ctx, collection := Connect()
	defer client.Disconnect(ctx)

	filter := bson.M{"short": bson.M{"$regex": "^" + short}}
	fmt.Println(filter)

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
