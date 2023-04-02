package db

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TIMEOUT = 10 * time.Second

var (
	mongoClient          *mongo.Client
	mongoConnectionError error
	mongoExec            sync.Once
)

func MongoClient() (*mongo.Client, error) {
	mongoExec.Do(func() {
		var (
			MONGODB_URI      = os.Getenv("MONGODB_URI")
			MONGODB_USERNAME = os.Getenv("MONGODB_USERNAME")
			MONGODB_PASSWORD = os.Getenv("MONGODB_PASSWORD")
		)
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT)
		mc, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI).SetAuth(options.Credential{
			Username: MONGODB_USERNAME,
			Password: MONGODB_PASSWORD,
		}))
		if err != nil {
			mongoConnectionError = err
			return
		}

		if errPing := mc.Ping(ctx, nil); errPing != nil {
			mongoConnectionError = errPing
		}

		mongoClient = mc
	})
	return mongoClient, mongoConnectionError
}

func FileCollection() *mongo.Collection {
	return mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("file")
}

func UserCollection() *mongo.Collection {
	return mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("user")
}
