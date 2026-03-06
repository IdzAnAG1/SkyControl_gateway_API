package service

import (
	"context"
	authv1 "sc_gateway/api/skycontrol/generated/proto/auth/v1"
)

type AuthService struct {
	authv1.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(context.Context, *authv1.LoginRequest) (*authv1.LoginResponse, error) {

	return &authv1.LoginResponse{
		Token:  "",
		UserId: 0,
	}, nil
}

func (a *AuthService) Register(context.Context, *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {

	return &authv1.RegisterResponse{
		UserId: 0,
		Token:  "",
	}, nil
}
