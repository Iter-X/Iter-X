package server

import (
	"net"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/server/interceptor"
	"github.com/iter-x/iter-x/internal/service"
)

type GRPCServer struct {
	*grpc.Server
	network string
	addr    string
	timeout time.Duration
	logger  *zap.SugaredLogger
}

func NewGRPCServer(
	c *conf.Server_GRPC,
	i18nBundle *i18n.Bundle,
	auth *service.Auth,
	trip *service.Trip,
	poi *service.PointsOfInterestService,
	geo *service.GeoService,
	logger *zap.SugaredLogger,
) *GRPCServer {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.I18n(i18nBundle),
			interceptor.TokenValidation(auth),
			interceptor.ValidateX(),
		),
	)
	authV1.RegisterAuthServiceServer(server, auth)
	tripV1.RegisterTripServiceServer(server, trip)
	poiV1.RegisterPointsOfInterestServiceServer(server, poi)
	geoV1.RegisterGeoServiceServer(server, geo)
	interceptor.LoadMethodOptions(server)
	return &GRPCServer{
		Server:  server,
		network: c.Network,
		addr:    c.Addr,
		timeout: c.Timeout.AsDuration(),
		logger:  logger.Named("grpc"),
	}
}

func (s *GRPCServer) Start() error {
	listen, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	s.logger.Infof("grpc server listening on %s", listen.Addr().String())
	return s.Serve(listen)
}
