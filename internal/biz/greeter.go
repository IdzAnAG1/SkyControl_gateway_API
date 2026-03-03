package biz

import (
	"context"
	v1 "sc_gateway/api/skycontrol/common"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrorServiceNotFound is service not found.
	ErrorServiceNotFound = errors.NotFound(v1.ErrorReason_SERVICE_NOT_FOUND.String(), "service not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

type HealthUsecase struct{}

// NewGreeterUsecase new a Greeter usecase.
func NewHealthUsecase() *HealthUsecase {
	return &HealthUsecase{}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *HealthUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	log.Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
