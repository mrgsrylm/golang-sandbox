package main

import (
	"context"
	"fmt"
	"golang_redis/config"
	"log"
	"time"

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

	// perform basic diagnostic to check if the connection is work
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis connection was refused")
	}

	fmt.Println(status)

	go pingPeriodically(ctx, rdb)

	// wait until program stopping
	select {}
}

// if you want use goroutine
func pingPeriodically(ctx context.Context, rdb *redis.Client) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping ping to redis server.")
			return
		case <-time.Tick(10 * time.Second):
			status, err := rdb.Ping(ctx).Result()
			if err != nil {
				log.Println("redis server connection was refused:", err)
				continue
			}
			fmt.Println(status)
		}
	}
}
