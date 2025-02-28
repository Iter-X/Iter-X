package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/impl/ent"
)

type RefreshToken interface {
	Base[*ent.RefreshToken, *do.RefreshToken]
}
