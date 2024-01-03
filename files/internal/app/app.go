package app

import (
	"context"
	"fmt"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/config"
	grpcServer "github.com/blazee5/cloud-drive/files/internal/grpc"
	"github.com/blazee5/cloud-drive/files/internal/service"
	"github.com/blazee5/cloud-drive/files/internal/storage"
	"github.com/blazee5/cloud-drive/files/lib/db/aws"
	"github.com/blazee5/cloud-drive/files/lib/db/postgres"
	"github.com/blazee5/cloud-drive/files/lib/logger"
	"github.com/blazee5/cloud-drive/files/lib/tracer"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger()

	ctx, cancel := context.WithCancel(context.Background())
	awsClient := aws.NewAWSClient(cfg)
	db := postgres.NewPgxConn(ctx, cfg)
	trace := tracer.InitTracer("files microservice")

	storages := storage.NewStorage(db, awsClient, trace.Tracer)
	services := service.NewFileService(log, storages, trace.Tracer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.HTTPServer.Port))
	if err != nil {
		log.Info("failed to listen: %v", err)
	}
	stats := otelgrpc.NewServerHandler(
		otelgrpc.WithTracerProvider(trace.Provider),
		otelgrpc.WithPropagators(propagation.TraceContext{}),
	)

	s := grpc.NewServer(grpc.StatsHandler(stats))
	pb.RegisterFileServiceServer(s, grpcServer.NewServer(log, services, trace.Tracer))

	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Infof("error while listen server: %s", err)
		}
	}()

	defer log.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	defer cancel()
	s.GracefulStop()
	db.Close()
}
