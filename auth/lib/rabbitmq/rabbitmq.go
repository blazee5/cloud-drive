package rabbitmq

import (
	"fmt"
	"github.com/blazee5/cloud-drive/auth/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitMQConn(cfg *config.Config) *amqp.Connection {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(connAddr)

	if err != nil {
		log.Fatalf("error connect to rabbitmq: %v", err)
	}

	return conn
}

func NewChannelConn(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	return ch, nil
}

func NewQueueConn(ch *amqp.Channel, cfg *config.Config) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		cfg.RabbitMQ.Queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &q, nil
}
