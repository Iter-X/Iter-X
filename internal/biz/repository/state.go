package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type State[T *ent.State, R *do.State] interface {
	BaseRepo[T, R]
}

type StateRepo = State[*ent.State, *do.State]
