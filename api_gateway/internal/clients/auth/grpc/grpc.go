package grpc

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func NewAuthServiceClient(log *zap.SugaredLogger) pb.AuthServiceClient {
	timeout, err := time.ParseDuration(os.Getenv("CLIENT_TIMEOUT"))

	if err != nil {
		log.Fatalf("error while parse timeout: %v", err)
	}

	retries, err := strconv.Atoi(os.Getenv("RETRIES_COUNT"))

	if err != nil {
		log.Fatalf("error while parse retries count: %v", err)
	}

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retries)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	cc, err := grpc.Dial(os.Getenv("AUTH_SVC_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)

	if err != nil {
		log.Fatalf("error while connect to models client: %s", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func InterceptorLogger() grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		slog.Log(ctx, slog.Level(lvl), msg, fields)
	})
}
