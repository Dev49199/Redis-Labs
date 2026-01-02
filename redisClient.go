package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

