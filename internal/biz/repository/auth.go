package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/bo"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type SmsCoder interface {
	GetSmsCode(ctx context.Context, params *bo.SendSmsConfigParams) (*bo.SmsCodeItem, error)

	VerifySmsCode(ctx context.Context, params *bo.VerifySmsCodeParams) error
}

type Auth[T *ent.User, R *do.User] interface {
	BaseRepo[T, R]

	SmsCoder

	FindByEmail(ctx context.Context, email string) (*do.User, error)

	FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error)

	Create(ctx context.Context, user *do.User) (*do.User, error)

	Update(ctx context.Context, user *do.User) (*do.User, error)

	GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error)

	GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error)

	SaveRefreshToken(ctx context.Context, val *do.RefreshToken) error

	UpdateRefreshToken(ctx context.Context, val *do.RefreshToken) error

	FindByPhone(ctx context.Context, phone string) (*do.User, error)
}

type AuthRepo = Auth[*ent.User, *do.User]
