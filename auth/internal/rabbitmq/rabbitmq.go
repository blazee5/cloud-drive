package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github.com/blazee5/cloud-drive/auth/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
)

func NewRabbitMQConn(cfg *config.Config) *amqp.Connection {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)

	conn, err := amqp.Dial(connAddr)

	if err != nil {
		log.Fatalf("error connect to rabbitmq: %v", err)
	}

	return conn
}

func NewChannelConn(conn *amqp.Connection, log *zap.SugaredLogger) (*amqp.Channel, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Infof("failed to create a channel in rabbitmq: %v", err)
		return nil, err
	}

	return ch, nil
}

func NewQueueConn(ch *amqp.Channel, log *zap.SugaredLogger) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Infof("failed to declare a queue: %v", err)
		return nil, err
	}

	return &q, nil
}

func PublishMessage(ctx context.Context, message string, ch *amqp.Channel, q *amqp.Queue) error {
	err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		return err
	}

	return nil
}

func NewConsumer(ctx context.Context, ch *amqp.Channel, q *amqp.Queue, consumeName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(q.Name, consumeName, false, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	return msgs, errors.New("")
}
