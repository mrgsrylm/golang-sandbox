package main

import (
	"context"
	"fmt"
	"golang-redis/config"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadRedisCfg()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	defer rdb.Close()

	// Set "FOO" to be associated with "BAR"
	rdb.Set(ctx, "FOO", "BAR", 0)
	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	// Update "FOO" to be associated with 5
	rdb.Set(ctx, "FOO", 5, 0)
	intResult, err := rdb.Get(ctx, "FOO").Int()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %d\n", intResult)
	}
}
