package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Url    string
	DBName string
	Ctx    context.Context
}

func GetMongoConnetion(config Config) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Url))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(config.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(config.Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client

}
