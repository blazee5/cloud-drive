package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/api_gateway/internal/clients/auth/grpc"
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"github.com/blazee5/cloud-drive/api_gateway/lib/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Service struct {
	log    *zap.SugaredLogger
	api    pb.AuthServiceClient
	tracer trace.Tracer
}

func NewService(log *zap.SugaredLogger, tracer *tracer.JaegerTracing) *Service {
	return &Service{
		log:    log,
		api:    grpc.NewAuthServiceClient(log, tracer),
		tracer: tracer.Tracer,
	}
}

func (s *Service) SignUp(ctx context.Context, input domain.SignUpRequest) (string, error) {
	ctx, span := s.tracer.Start(ctx, "authService.SignUp")
	defer span.End()

	res, err := s.api.SignUp(ctx, &pb.SignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}

func (s *Service) SignIn(ctx context.Context, input domain.SignInRequest) (string, error) {
	ctx, span := s.tracer.Start(ctx, "authService.SignIn")
	defer span.End()

	res, err := s.api.SignIn(ctx, &pb.SignInRequest{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return "", err
	}

	return res.GetToken(), nil
}

func (s *Service) ActivateAccount(ctx context.Context, code string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "authService.ActivateAccount")
	defer span.End()

	res, err := s.api.ValidateAccount(ctx, &pb.ValidateAccountRequest{
		Code: code,
	})

	if err != nil {
		return "", err
	}

	return res.GetStatus(), nil
}
