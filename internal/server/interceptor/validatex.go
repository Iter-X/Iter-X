package interceptor

import (
	"context"
	"errors"

	errors1 "github.com/protoc-gen/protoc-gen-go-errors/errors"
	"google.golang.org/grpc"

	"github.com/iter-x/iter-x/internal/common/xerr"
)

type validateX interface {
	Validate(context.Context) error
}

func ValidateX() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(validateX); ok {
			if err := v.Validate(ctx); err != nil {
				var xErr *errors1.Error
				if errors.As(err, &xErr) {
					return nil, err
				}
				return nil, xerr.New(400, "INVALID_PARAMETERS", err.Error())
			}
		}
		return handler(ctx, req)
	}
}
