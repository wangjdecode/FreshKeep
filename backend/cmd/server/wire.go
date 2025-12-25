//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/freshkeep/backend/internal/conf"
	"github.com/freshkeep/backend/internal/data"
	"github.com/freshkeep/backend/internal/service"
)

// wireApp init kratos application.
func wireApp(srv *conf.Server, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		service.ProviderSet,
		NewHTTPServer,
		grpc.NewServer,
		newApp,
	))
}

