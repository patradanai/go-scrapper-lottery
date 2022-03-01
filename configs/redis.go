package configs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ClientRedis *redis.Client

func ConnectionRedis() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	ClientRedis = client

	fmt.Printf("Redis Connected : %v \n", pong)
	return nil
}
