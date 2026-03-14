package service

import (
	"context"
	authv1 "sc_gateway/api/skycontrol/generated/proto/auth/v1"
	"sc_gateway/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	authv1.UnimplementedAuthServer
	authClient authv1.AuthClient
	logger     log.Logger
}

func NewAuthService(data *data.Data, logger log.Logger) *AuthService {
	return &AuthService{
		authClient: data.AuthClient,
		logger:     logger,
	}
}

func (a *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	err := a.logger.Log(log.LevelInfo, "msg", "[gRPCClient] Authenticating user...")
	if err != nil {
		return nil, err
	}
	return a.authClient.Login(ctx, req)
}

func (a *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	err := a.logger.Log(log.LevelInfo, "msg", "[gRPCClient] Authenticating user...")
	if err != nil {
		return nil, err
	}
	return a.authClient.Register(ctx, req)
}
