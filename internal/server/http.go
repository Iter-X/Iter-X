package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	userV1 "github.com/iter-x/iter-x/internal/api/user/v1"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/conf"
)

type HTTPServer struct {
	env    conf.Environment
	cfg    *conf.Server_HTTP
	mux    *runtime.ServeMux
	logger *zap.SugaredLogger
}

func NewHTTPServer(env conf.Environment, c *conf.Server_HTTP, logger *zap.SugaredLogger) *HTTPServer {
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
		env:    env,
		cfg:    c,
		mux:    mux,
		logger: logger.Named("http"),
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	if err := registerDoc(s.env, s.mux); err != nil {
		return err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	for _, fn := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		authV1.RegisterAuthServiceHandlerFromEndpoint,
		tripV1.RegisterTripServiceHandlerFromEndpoint,
		poiV1.RegisterPointsOfInterestServiceHandlerFromEndpoint,
		geoV1.RegisterGeoServiceHandlerFromEndpoint,
		userV1.RegisterUserServiceHandlerFromEndpoint,
	} {
		err := fn(ctx, s.mux, s.cfg.GrpcAddr, opts)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("http server listening on %s", s.cfg.Addr)
	return http.ListenAndServe(s.cfg.Addr, s.mux)
}

func registerDoc(env conf.Environment, mux *runtime.ServeMux) error {
	if env != conf.Environment_DEV && env != conf.Environment_TEST {
		return nil
	}
	err := mux.HandlePath(http.MethodGet, "/doc/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/doc/", http.FileServer(http.Dir("./swagger"))).ServeHTTP(w, r)
	})
	if err != nil {
		return err
	}
	err = mux.HandlePath(http.MethodPost, "/doc/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/doc/", http.FileServer(http.Dir("./swagger"))).ServeHTTP(w, r)
	})
	if err != nil {
		return err
	}
	err = mux.HandlePath(http.MethodGet, "/dbviewer/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/dbviewer/", http.FileServer(http.Dir("./dbviewer"))).ServeHTTP(w, r)
	})
	if err != nil {
		return err
	}
	err = mux.HandlePath(http.MethodPost, "/dbviewer/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/dbviewer/", http.FileServer(http.Dir("./dbviewer"))).ServeHTTP(w, r)
	})
	if err != nil {
		return err
	}
	return nil
}
