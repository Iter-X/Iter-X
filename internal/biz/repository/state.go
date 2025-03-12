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

	// ListStates 列出州/省，可选按国家过滤
	ListStates(ctx context.Context, params *bo.ListStatesParams) ([]*do.State, *bo.PaginationResult, error)
}

type StateRepo = State[*ent.State, *do.State]
