package interceptor

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type TokenValidator interface {
	ValidateToken(ctx context.Context, token string) (jwt.Claims, error)
}

var (
	skipTokenMap = map[string]bool{} // ReadOnly
)

func LoadMethodOptions(server *grpc.Server) {
	sds, _ := grpcreflect.LoadServiceDescriptors(server)
	for _, sd := range sds {
		for _, md := range sd.GetMethods() {
			methodName := fmt.Sprintf("/%s/%s", sd.GetFullyQualifiedName(), md.GetName())
			ext := proto.GetExtension(md.GetMethodOptions(), cnst.E_SkipToken)
			if ext == true {
				skipTokenMap[methodName] = true
			}
		}
	}
}

func TokenValidation(validator TokenValidator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if skipTokenMap[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, xerr.ErrorUnauthorized()
		}

		tokens := md.Get(cnst.HttpHeaderAuthorization)
		if len(tokens) == 0 {
			return nil, xerr.ErrorUnauthorized()
		}
		var (
			claims jwt.Claims
			err    error
		)
		claims, err = validator.ValidateToken(ctx, tokens[0])
		if err != nil {
			if xerr.As(err) != nil {
				return nil, err
			}
			return nil, xerr.ErrorUnauthorized()
		}
		ctx = context.WithValue(ctx, cnst.CtxKeyClaims, claims)
		return handler(ctx, req)
	}
}
