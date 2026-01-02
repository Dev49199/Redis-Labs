package main

import (
        "fmt"
        "github.com/redis/go-redis/v9"
)

func saveUser(rdb *redis.Client) {
	key := "user:1"

	rdb.HSet(ctx, key, map[string]interface{}{
		"name": "Dev",
		"age":  27,
		"role": "engineer",
	})
}

func incrementAge(rdb *redis.Client) {
	rdb.HIncrBy(ctx, "user:1", "age", 1)
}

func getUser(rdb *redis.Client) {
	data, _ := rdb.HGetAll(ctx, "user:1").Result()
	fmt.Println(data)
}

func main(){
rdb := redisClient()
saveUser(rdb)
incrementAge(rdb)
getUser(rdb)
}
