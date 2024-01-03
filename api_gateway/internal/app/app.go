package app

import (
	"context"
	"github.com/blazee5/cloud-drive/api_gateway/internal/server"
	"github.com/blazee5/cloud-drive/api_gateway/lib/tracer"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *zap.SugaredLogger) {
	trace := tracer.InitTracer("api gateway")
	srv := server.NewServer(log, trace)
	go func() {
		if err := srv.Run(srv.InitRoutes()); err != nil {
			log.Fatalf("Error while start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Infof("Error occured on server shutting down: %v", err)
	}
}
