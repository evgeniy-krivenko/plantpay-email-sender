package main

import (
	"fmt"
	"github.com/evgeniy-krivenko/plantpay-email-sender/internal/handler"
	"github.com/evgeniy-krivenko/plantpay-email-sender/pkg/mailer"
	"github.com/evgeniy-krivenko/plantpay-email-sender/pkg/qpye"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error load env file: %s", err.Error())
	}
	conn, err := amqp.Dial(os.Getenv("AMQP"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Println("Connected to RMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	emailPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		logrus.Fatalf("error port %s", err.Error())
	}
	// init mailer
	ml := mailer.New(mailer.Config{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     emailPort,
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Sender:   os.Getenv("EMAIL_SENDER"),
		Timeout:  5 * time.Second,
	})
	// init handler
	h := handler.NewHandler(ml)
	// init amqp server
	qr := qpye.NewServer(conn, ch, h.InitRoutes())
	// start listen
	qr.ListenQueue("order_queue")
}
