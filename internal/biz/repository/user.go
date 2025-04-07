package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/data/ent"

	"github.com/iter-x/iter-x/internal/biz/do"
)

type User[T *ent.User, R *do.User] interface {
	BaseRepo[T, R]

	FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error)
	Update(ctx context.Context, user *do.User) (*do.User, error)
	GetUserPreference(ctx context.Context, userId uuid.UUID) (*do.UserPreference, error)
	CreateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error
	UpdateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error
}

type UserRepo = User[*ent.User, *do.User]
