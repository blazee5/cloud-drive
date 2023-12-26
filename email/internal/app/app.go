package app

import (
	"context"
	consumer "github.com/blazee5/cloud-drive/email/internal/email/handler/rabbitmq"
	"github.com/blazee5/cloud-drive/email/internal/email/service"
	"github.com/blazee5/cloud-drive/email/lib/rabbitmq"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *zap.SugaredLogger) {
	conn := rabbitmq.NewRabbitMQConn()

	ch, err := rabbitmq.NewChannelConn(conn)

	if err != nil {
		log.Fatalf("failed to create a channel in rabbitmq: %v", err)
	}

	_, err = rabbitmq.NewQueueConn(ch)

	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	emailService := service.NewService(log)
	c := consumer.NewConsumer(log, emailService)

	log.Infof("email service starting...")

	err = c.ConsumeQueue(ctx, ch)

	if err != nil {
		log.Infof("error: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err = ch.Close()

	if err != nil {
		log.Infof("error while close channel: %v", err)
	}

	err = conn.Close()

	if err != nil {
		log.Infof("error while close rabbitmq conn: %v", err)

	}
}
