package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Auth[T *ent.User, R *do.User] interface {
	BaseRepo[T, R]

	FindByEmail(ctx context.Context, email string) (*do.User, error)

	FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error)

	Create(ctx context.Context, user *do.User) (*do.User, error)

	Update(ctx context.Context, user *do.User) (*do.User, error)

	GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error)

	GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error)

	SaveRefreshToken(ctx context.Context, val *do.RefreshToken) error

	UpdateRefreshToken(ctx context.Context, val *do.RefreshToken) error

	FindByPhone(ctx context.Context, phone string) (*do.User, error)

	GetUserPreference(ctx context.Context, userId uuid.UUID) (*do.UserPreference, error)

	CreateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error

	UpdateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error
}

type AuthRepo = Auth[*ent.User, *do.User]
