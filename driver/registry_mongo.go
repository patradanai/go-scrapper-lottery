package driver

import (
	"context"
	"lottery-web-scrapping/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMongo *mongo.Client

func ConnectionMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.LoadEnv("MONGO_DB")))
	if err != nil {
		return err
	}

	ClientMongo = client
	return nil
}
