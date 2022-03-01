package configs

import (
	"context"
	"lottery-web-scrapping/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientMongo *mongo.Client

func ConnectionMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.LoadEnv("MONGO_DB")))
	if err != nil {
		return err
	}

	clientMongo = client
	return nil
}
