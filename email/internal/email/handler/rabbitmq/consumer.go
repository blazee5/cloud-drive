package rabbitmq

import (
	"context"
	"github.com/blazee5/cloud-drive/email/internal/email"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Consumer struct {
	log     *zap.SugaredLogger
	service email.Service
}

func NewConsumer(log *zap.SugaredLogger, service email.Service) *Consumer {
	return &Consumer{log: log, service: service}
}

func (c *Consumer) RunConsumer(ctx context.Context, ch <-chan amqp.Delivery) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-ch:
			if !ok {
				c.log.Infof("channel is closed")
			}

			if err := c.service.SendEmail(string(msg.Body)); err != nil {
				c.log.Infof("error while send email: %v", err)
				continue
			}
		}
	}
}
