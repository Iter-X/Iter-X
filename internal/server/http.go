package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/conf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HTTPServer struct {
	cfg    *conf.Server_HTTP
	mux    *runtime.ServeMux
	logger *zap.SugaredLogger
}

func NewHTTPServer(c *conf.Server_HTTP, logger *zap.SugaredLogger) *HTTPServer {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case cnst.HttpHeaderLang, cnst.HttpHeaderAcceptLang:
				return key, true
			default:
				return key, false
			}
		}),
	)
	return &HTTPServer{
		cfg:    c,
		mux:    mux,
		logger: logger.Named("http"),
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	for _, fn := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		authV1.RegisterAuthServiceHandlerFromEndpoint,
		tripV1.RegisterTripServiceHandlerFromEndpoint,
		poiV1.RegisterPointsOfInterestServiceHandlerFromEndpoint,
	} {
		err := fn(ctx, s.mux, s.cfg.GrpcAddr, opts)
		if err != nil {
			return err
		}
	}
	s.logger.Infof("http server listening on %s", s.cfg.Addr)
	return http.ListenAndServe(s.cfg.Addr, s.mux)
}
