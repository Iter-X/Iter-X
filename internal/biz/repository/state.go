package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type State[T *ent.State, R *do.State] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error)

	// ListStates lists states/provinces, optionally filtered by country
	ListStates(ctx context.Context, params *bo.ListStatesParams) ([]*do.State, int64, error)
}

type StateRepo = State[*ent.State, *do.State]
