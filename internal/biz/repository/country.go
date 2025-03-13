package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Country[T *ent.Country, R *do.Country] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error)

	// ListCountries lists countries, optionally filtered by continent
	ListCountries(ctx context.Context, params *bo.ListCountriesParams) ([]*do.Country, int64, error)
}

type CountryRepo = Country[*ent.Country, *do.Country]
