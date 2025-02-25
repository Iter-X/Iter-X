// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/i18n"
	"github.com/iter-x/iter-x/internal/repo"
	"github.com/iter-x/iter-x/internal/server"
	"github.com/iter-x/iter-x/internal/service"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func wireApp(i18nCfg *conf.I18N, grpcCfg *conf.Server_GRPC, httpCfg *conf.Server_HTTP, data *conf.Data, authCfg *conf.Auth, agentCfg *conf.Agent, logger *zap.SugaredLogger) (*App, func(), error) {
	bundle := i18n.New(i18nCfg)
	client, cleanup, err := repo.NewConnection(data, logger)
	if err != nil {
		return nil, nil, err
	}
	auth := repo.NewAuth(client, logger)
	bizAuth := biz.NewAuth(authCfg, auth, logger)
	serviceAuth := service.NewAuth(bizAuth)
	trip := repo.NewTrip(client, logger)
	bizTrip := biz.NewTrip(trip, logger)
	serviceTrip := service.NewTrip(bizTrip)
	grpcServer := server.NewGRPCServer(grpcCfg, bundle, serviceAuth, serviceTrip, logger)
	httpServer := server.NewHTTPServer(httpCfg, logger)
	app := newApp(grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
