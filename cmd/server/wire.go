//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/ai"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/impl"
	"github.com/iter-x/iter-x/internal/helper/i18n"
	"github.com/iter-x/iter-x/internal/server"
	"github.com/iter-x/iter-x/internal/service"
	"go.uber.org/zap"
)

func wireApp(env conf.Environment, i18nCfg *conf.I18N, grpcCfg *conf.Server_GRPC, httpCfg *conf.Server_HTTP, d *conf.Data, authCfg *conf.Auth, storageCfg *conf.Storage, logger *zap.SugaredLogger) (*App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		impl.ProviderSet,
		biz.ProviderSet,
		i18n.ProviderSet,
		ai.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp))
}
