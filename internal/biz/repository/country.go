package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Country[T *ent.Country, R *do.Country] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error)
}

type CountryRepo = Country[*ent.Country, *do.Country]
