package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type City[T *ent.City, R *do.City] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error)
}

type CityRepo = City[*ent.City, *do.City]
