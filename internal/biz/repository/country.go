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

	// ListCountries 列出国家，可选按大洲过滤
	ListCountries(ctx context.Context, params *bo.ListCountriesParams) ([]*do.Country, *bo.PaginationResult, error)
}

type CountryRepo = Country[*ent.Country, *do.Country]
