package main
import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)
func apiHit(rdb *redis.Client, userID string) int64 {
	key := "api_hits:" + userID

	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	if count == 1 {
		rdb.Expire(ctx, key, time.Minute)
	}

	return count
}

func main() {
	rdb := redisClient()

	for i := 0; i < 5; i++ {
		fmt.Println("Hits:", apiHit(rdb, "user1"))
	}
}
