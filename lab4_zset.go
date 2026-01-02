package main

import (
        "fmt"
        "github.com/redis/go-redis/v9"
)

func addScore(rdb *redis.Client, user string, score float64) {
	rdb.ZAdd(ctx, "leaderboard", redis.Z{
		Score:  score,
		Member: user,
	})
}

func topPlayers(rdb *redis.Client) {
	players, _ := rdb.ZRevRangeWithScores(ctx, "leaderboard", 0, 2).Result()
	for _, p := range players {
		fmt.Println(p.Member, p.Score)
	}
}

func main(){
rdb := redisClient()
addScore(rdb,"user_1",100.34)
addScore(rdb, "user_2",101.31)
addScore(rdb, "user_3",99)
addScore(rdb,"user_4",97)
topPlayers(rdb)
}
