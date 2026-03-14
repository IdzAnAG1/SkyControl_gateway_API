//go:build wireinject

// go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"sc_gateway/internal/conf"
	"sc_gateway/internal/data"
	"sc_gateway/internal/server"
	"sc_gateway/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
