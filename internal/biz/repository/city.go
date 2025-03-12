package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type City[T *ent.City, R *do.City] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error)

	// ListCities 列出城市，可选按州/省过滤
	ListCities(ctx context.Context, params *bo.ListCitiesParams) ([]*do.City, *bo.PaginationResult, error)
}

type CityRepo = City[*ent.City, *do.City]
