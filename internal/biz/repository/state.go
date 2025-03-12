package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type State[T *ent.State, R *do.State] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error)
}

type StateRepo = State[*ent.State, *do.State]
