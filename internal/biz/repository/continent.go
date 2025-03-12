package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Continent[T *ent.Continent, R *do.Continent] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error)
}

type ContinentRepo = Continent[*ent.Continent, *do.Continent]
