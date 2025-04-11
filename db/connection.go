package db

import (
	"context"
	"sync"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbname = "go_todo"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
	ctx            context.Context
)

func getClient() (*mongo.Client, context.Context) {
	clientOnce.Do(func() {
		var err error
		clientInstance, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = clientInstance.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	})

	return clientInstance, context.Background()
}

func CloseClient() {
	if clientInstance != nil {
		_ = clientInstance.Disconnect(ctx)
	}
}
