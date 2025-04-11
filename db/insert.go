package db

import "go.mongodb.org/mongo-driver/bson/primitive"

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getClient()

	c := client.Database(dbname).Collection(collection)

	resp, err := c.InsertOne(ctx, data)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return resp.InsertedID.(primitive.ObjectID), nil
}
