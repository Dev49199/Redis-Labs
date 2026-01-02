package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var streamCtx = context.Background()

const (
	streamName  = "orders_stream"
	groupName   = "order_processors"
	consumerName = "consumer-1"
)

// ---------------- Producer ----------------
func publishOrder(rdb *redis.Client, orderID string) error {
	_, err := rdb.XAdd(streamCtx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"order_id": orderID,
			"status":   "CREATED",
			"time":     time.Now().Format(time.RFC3339),
		},
	}).Result()

	return err
}

// ---------------- Consumer ----------------
func createConsumerGroup(rdb *redis.Client) {
	err := rdb.XGroupCreateMkStream(
		streamCtx,
		streamName,
		groupName,
		"$",
	).Err()

	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}
}

func consumeOrders(rdb *redis.Client) {
	for {
		streams, err := rdb.XReadGroup(streamCtx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamName, ">"},
			Count:    1,
			Block:    5 * time.Second,
		}).Result()

		if err == redis.Nil {
			continue
		}
		if err != nil {
			panic(err)
		}

		for _, stream := range streams {
			for _, msg := range stream.Messages {
				fmt.Printf("📦 Processing order %v\n", msg.Values["order_id"])

				// Simulate work
				time.Sleep(1 * time.Second)

				// Acknowledge message
				rdb.XAck(streamCtx, streamName, groupName, msg.ID)
				fmt.Printf("✅ Order %v processed\n", msg.Values["order_id"])
			}
		}
	}
}

// ---------------- Main ----------------
func main() {
	rdb := redisClient()
	createConsumerGroup(rdb)

	// Producer (simulate incoming orders)
	go func() {
		for i := 1; i <= 5; i++ {
			orderID := fmt.Sprintf("order-%d", i)
			err := publishOrder(rdb, orderID)
			if err != nil {
				panic(err)
			}
			fmt.Println("📝 Published:", orderID)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Consumer
	consumeOrders(rdb)
}

