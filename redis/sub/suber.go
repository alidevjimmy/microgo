package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pubsub := rdb.Subscribe(ctx, "mychannel1")
	defer pubsub.Close()

	forever := make(chan bool)

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(msg)
		}
	}()

	fmt.Println("Subscriber in waiting for messages!")
	<-forever
}
