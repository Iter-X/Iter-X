package interceptor

import (
	"context"
	"errors"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	xerr "github.com/protoc-gen/protoc-gen-go-errors/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func I18n(bundle *i18n.Bundle) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			var xErr *xerr.Error
			if !errors.As(err, &xErr) {
				return nil, err
			}
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, err
			}
			localize, err2 := i18n.NewLocalizer(bundle,
				getHeader(md, cnst.HttpHeaderLang), getHeader(md, cnst.HttpHeaderAcceptLang)).
				Localize(&i18n.LocalizeConfig{MessageID: xErr.GetReason()})
			if err2 == nil {
				xErr.Message = localize
			}
			return nil, err
		}
		return resp, nil
	}
}

func getHeader(md metadata.MD, key string) string {
	if md == nil {
		return ""
	}
	vals := md.Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}
