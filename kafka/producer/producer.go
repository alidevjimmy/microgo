package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/Shopify/sarama.v1"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, cfg)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushCommentToQueue(producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func main() {
	brokersUrl := []string{"127.0.0.1:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Close()

	topic := "comments"
	msg := []byte(os.Args[1])
	if err := PushCommentToQueue(producer, topic, msg); err != nil {
		log.Fatalln(err)
	}
}
