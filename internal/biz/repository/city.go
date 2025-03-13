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

	// ListCities lists cities, optionally filtered by state/province
	ListCities(ctx context.Context, params *bo.ListCitiesParams) ([]*do.City, int64, error)
}

type CityRepo = City[*ent.City, *do.City]
