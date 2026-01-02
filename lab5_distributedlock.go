package main

import (
        "fmt"
        "github.com/redis/go-redis/v9"
"time"
)

func acquireLock(rdb *redis.Client, key string) bool {
	ok, _ := rdb.SetNX(ctx, key, "worker1", 100*time.Second).Result()
	return ok
}

func main() {
	rdb := redisClient()

	if acquireLock(rdb, "lock:order:1") {
		fmt.Println("Processing order")
		time.Sleep(30 * time.Second)
		rdb.Del(ctx, "lock:order:1")
	} else {
		fmt.Println("Order already being processed")
	}
}
