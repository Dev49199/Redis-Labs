package main

import (
        "fmt"
        "github.com/redis/go-redis/v9"
)

func userOnline(rdb *redis.Client, userID string) {
	rdb.SAdd(ctx, "online_users", userID)
}

func userOffline(rdb *redis.Client, userID string) {
	rdb.SRem(ctx, "online_users", userID)
}

func isOnline(rdb *redis.Client, userID string) bool {
	ok, _ := rdb.SIsMember(ctx, "online_users", userID).Result()
	return ok
}


func main(){
rdb:=redisClient()
userOnline(rdb,"1241")
userOnline(rdb,"1242")
userOnline(rdb,"1243")

userOffline(rdb,"1242")

userOnline := isOnline(rdb,"1242")
fmt.Println("User 1242 online",userOnline)
}
