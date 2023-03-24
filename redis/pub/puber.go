package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Publish(ctx, "mychannel1", os.Args[1]).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("message published!")
}
