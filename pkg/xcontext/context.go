package xcontext

import (
	"context"

	"github.com/iter-x/iter-x/pkg/vobj"
)

type clientTypeContextKey struct{}

// WithClientType returns a copy of parent in which the client type value is set.
func WithClientType(parent context.Context, clientType string) context.Context {
	return context.WithValue(parent, clientTypeContextKey{}, vobj.BuildClientType(clientType))
}

// ClientTypeFrom returns the client type value stored in ctx, if any.
func ClientTypeFrom(ctx context.Context) (string, bool) {
	clientType, ok := ctx.Value(clientTypeContextKey{}).(vobj.ClientType)
	return clientType.String(), ok
}
