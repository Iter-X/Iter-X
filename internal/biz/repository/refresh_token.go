package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/impl/ent"
)

type RefreshToken[T *ent.RefreshToken, R *do.RefreshToken] interface {
	Base[T, R]
}

type RefreshTokenRepo = RefreshToken[*ent.RefreshToken, *do.RefreshToken]
