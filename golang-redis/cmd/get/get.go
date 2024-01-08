package main

import (
	"context"
	"fmt"
	"golang-redis/config"

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

	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("Key FOO not found in Redis cache")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	intValue, err := rdb.Get(ctx, "INT").Int()
	if err != nil {
		fmt.Println("Key INT not found in Redis cache")
	} else {
		fmt.Printf("INT has value %d\n", intValue)
	}

	var person Person
	err = rdb.HGetAll(ctx, "STRUCT").Scan(&person)
	if err != nil {
		fmt.Println("Key STRUCT not found in Redis cache")
	} else {
		fmt.Printf("STRUCT has value %+v\n", person)
	}

	result, err = rdb.Get(ctx, "BAZ").Result()
	if err != nil {
		fmt.Println("Key BAZ not found in Redis cache")
	} else {
		fmt.Printf("BAZ has value %s\n", result)
	}
}
