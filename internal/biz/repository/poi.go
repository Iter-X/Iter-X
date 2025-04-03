package repository

import (
	"context"
	"github.com/google/uuid"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type PointsOfInterest[T *ent.PointsOfInterest, R *do.PointsOfInterest] interface {
	BaseRepo[T, R]

	SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error)
	SearchPointsOfInterestByNamesFromES(ctx context.Context, names []string) ([]*do.PointsOfInterest, error)

	GetByCityNames(ctx context.Context, cityNames []string) ([]*do.PointsOfInterest, error)

	GetTopPOIsByCity(ctx context.Context, cityIds []int32, limit int) ([]*do.PointsOfInterest, error)

	GetByIds(ctx context.Context, ids []uuid.UUID) ([]*do.PointsOfInterest, error)
}

type PointsOfInterestRepo = PointsOfInterest[*ent.PointsOfInterest, *do.PointsOfInterest]
