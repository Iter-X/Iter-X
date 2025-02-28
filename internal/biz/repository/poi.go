package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type PointsOfInterest[T *ent.PointsOfInterest, R *do.PointsOfInterest] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error)
}

type PointsOfInterestRepo = PointsOfInterest[*ent.PointsOfInterest, *do.PointsOfInterest]
