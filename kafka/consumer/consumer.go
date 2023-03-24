package main

import (
	"fmt"
	"log"

	"gopkg.in/Shopify/sarama.v1"
)

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, cfg)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	brokersUrl := []string{"127.0.0.1:9092"}
	worker, err := ConnectConsumer(brokersUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer worker.Close()

	consumer, err := worker.ConsumePartition("comments", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln(err)
	}
	defer consumer.Close()

	forever := make(chan bool)
	go func() {
		for msg := range consumer.Messages() {
			fmt.Println(msg)
		}
	}()

	fmt.Println("Waiting for Message...")
	<-forever

}
