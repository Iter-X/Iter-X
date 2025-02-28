package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Token[T *ent.RefreshToken, R *do.RefreshToken] interface {
	BaseRepo[T, R]
}

type TokenRepo = Token[*ent.RefreshToken, *do.RefreshToken]
