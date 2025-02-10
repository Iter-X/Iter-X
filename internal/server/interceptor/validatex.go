package interceptor

import (
	"context"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"google.golang.org/grpc"
)

type validatex interface {
	Validate(context.Context) error
}

func Validatex() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(validatex); ok {
			if err := v.Validate(ctx); err != nil {
				return nil, xerr.New(400, "INVALID_PARAMETERS", err.Error())
			}
		}
		return handler(ctx, req)
	}
}
