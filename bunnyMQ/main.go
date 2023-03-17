package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Connect(c QeueConfig) (*amqp.Connection, error) {
	connPattern := "amqp://%s:%s@%s:%s/%s"
	if c.Username == "" {
		connPattern = "amqp://%s:%s@%s:%s"
	}

	clientUrl := fmt.Sprintf(connPattern, c.Username, c.Password, c.Host, c.Port, c.Vhost)
}