package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectMQ() (*amqp.Connection, error) {
	connPattern := "amqp://%s:%s@%s:%s/%s"

	clientUrl := fmt.Sprintf(connPattern, "c.Username", "c.Password", "c.Host", "c.Port", "c.Vhost")
	conn, err := amqp.Dial(clientUrl)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	return conn, err

}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Logger())

	rabbitMq, err := ConnectMQ()
	if err != nil {
		log.Println(err)
		return
	}
	defer rabbitMq.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	log.Fatal(e.Start(":8080"))
}
