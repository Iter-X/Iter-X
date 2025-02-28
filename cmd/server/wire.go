//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/i18n"
	"github.com/iter-x/iter-x/internal/impl"
	"github.com/iter-x/iter-x/internal/server"
	"github.com/iter-x/iter-x/internal/service"
	"go.uber.org/zap"
)

func wireApp(i18nCfg *conf.I18N, grpcCfg *conf.Server_GRPC, httpCfg *conf.Server_HTTP, data *conf.Data, authCfg *conf.Auth, agentCfg *conf.Agent, logger *zap.SugaredLogger) (*App, func(), error) {
	panic(wire.Build(
		i18n.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
		biz.ProviderSet,
		impl.ProviderSet,
		newApp))
}
