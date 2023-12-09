package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func NewRabbitMQConn() *amqp.Connection {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
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

func NewQueueConn(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE"),
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
