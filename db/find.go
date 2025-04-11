package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll(collection string, filter bson.M, documents any) error {
	client, ctx := getClient()

	c := client.Database(dbname).Collection(collection)
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, documents)
}

func FindByID(collection string, id string, document any) error {
	client, ctx := getClient()

	c := client.Database(dbname).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	return c.FindOne(ctx, filter).Decode(document)
}

func FindOne(collection string, filter bson.M, document any) error {
	client, ctx := getClient()

	c := client.Database(dbname).Collection(collection)

	return c.FindOne(ctx, filter).Decode(document)
}
