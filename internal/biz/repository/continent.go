package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Continent[T *ent.Continent, R *do.Continent] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error)

	// ListContinents 列出所有大洲
	ListContinents(ctx context.Context, params *bo.ListContinentsParams) ([]*do.Continent, *bo.PaginationResult, error)
}

type ContinentRepo = Continent[*ent.Continent, *do.Continent]
