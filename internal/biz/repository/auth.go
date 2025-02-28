package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/impl/ent"
)

type Auth[T *ent.User, R *do.User] interface {
	Base[T, R]

	FindByEmail(ctx context.Context, email string) (*do.User, error)

	FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error)

	Create(ctx context.Context, user *do.User) (*do.User, error)

	Update(ctx context.Context, user *do.User) (*do.User, error)

	GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error)

	GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error)

	SaveRefreshToken(ctx context.Context, val *do.RefreshToken) error

	UpdateRefreshToken(ctx context.Context, val *do.RefreshToken) error
}

type AuthRepo = Auth[*ent.User, *do.User]
