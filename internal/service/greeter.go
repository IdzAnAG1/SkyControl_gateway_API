package service

import (
	"context"
	"fmt"
	v1 "sc_gateway/api/skycontrol/viability"
	"sc_gateway/internal/service/variables"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type HealthService struct {
	v1.UnimplementedViabilityServer
	uptime time.Time
}

// NewHealthService new a greeter service.
func NewHealthService() *HealthService {
	return &HealthService{
		uptime: time.Now(),
	}
}

// Health implements checking service viability
func (s *HealthService) Health(context.Context, *emptypb.Empty) (*v1.HealthReply, error) {
	/*
		todo check services
	*/
	return &v1.HealthReply{
		GatewayStatus: fmt.Sprintf(variables.ServiceIsUp, "Gateway"),
		GatewayUptime: s.uptime.Format("2006-01-02 15:04:05"),
	}, nil
}
