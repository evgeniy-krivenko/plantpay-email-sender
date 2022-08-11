package qpye

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Event struct {
	Pattern string          `json:"pattern"`
	Data    json.RawMessage `json:"data"`
}

type Handler func(d *json.RawMessage)

type Server struct {
	*Router
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (qr *Server) ListenQueue(q string) {
	msgs, _ := qr.ch.Consume(q,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var e Event
			err := json.Unmarshal(d.Body, &e)
			if err != nil {
				fmt.Printf(err.Error())
			}

			if cb, ok := qr.Routes[e.Pattern]; ok {
				go cb(&e.Data)
			}
		}
	}()
	fmt.Println("Connected to RabbitMQ instance")
	<-forever
}

func NewServer(conn *amqp.Connection, ch *amqp.Channel, r *Router) *Server {
	return &Server{
		Router: r,
		conn:   conn,
		ch:     ch,
	}
}
