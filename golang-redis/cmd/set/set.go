package main

import (
	"context"
	"fmt"
	"golang-redis/config"
	"time"

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

	_, err := rdb.Set(ctx, "FOO", "BAR", 0).Result()
	if err != nil {
		fmt.Println("Failed to add FOO <> BAR key-value pair")
		return
	}
	rdb.Set(ctx, "INT", 5, 0)
	rdb.Set(ctx, "FLOAT", 5.5, 0)
	rdb.Set(ctx, "EXPIRING", 15, 30*time.Minute)
	rdb.Set(ctx, "LIST", []string{"Hello"}, 0)

	rdb.HSet(ctx, "STRUCT", Person{"John Doe", 15})
}
