package main

import (
	"context"
	"fmt"
	"golang_redis/config"

	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func main() {
	ctx := context.Background()
	cfg := config.LoadRedisCfg()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	defer rdb.Close()

	rdb.Set(ctx, "FOO", "BAR", 0)
	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	// Deleting the key "FOO" and its associated value
	rdb.Del(ctx, "FOO")
	result, err = rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}
}
