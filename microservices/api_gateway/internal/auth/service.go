package auth

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/domain"
)

type Service interface {
	SignUp(ctx context.Context, input domain.SignUpRequest) (string, error)
	SignIn(ctx context.Context, input domain.SignInRequest) (string, error)
	ValidateUser(ctx context.Context, token string) (string, error)
}
