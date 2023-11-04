package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoClient(connectionCtx context.Context, opts *options.ClientOptions) *mongo.Client {
	client, err := mongo.Connect(connectionCtx, opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
