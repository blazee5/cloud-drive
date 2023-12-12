package rabbitmq

import (
	"context"
	"github.com/blazee5/cloud-drive/auth/internal/config"
	"github.com/blazee5/cloud-drive/auth/lib/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Producer struct {
	log   *zap.SugaredLogger
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue *amqp.Queue
}

func NewProducer(log *zap.SugaredLogger, conn *amqp.Connection) *Producer {
	return &Producer{log: log, conn: conn}
}

func (p *Producer) InitProducer(cfg *config.Config) error {
	ch, err := rabbitmq.NewChannelConn(p.conn)

	if err != nil {
		return err
	}

	q, err := rabbitmq.NewQueueConn(ch, cfg)

	if err != nil {
		return err
	}

	p.ch = ch
	p.queue = q

	return nil
}

func (p *Producer) PublishMessage(ctx context.Context, message string) error {
	err := p.ch.PublishWithContext(ctx,
		"",
		p.queue.Name,
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
